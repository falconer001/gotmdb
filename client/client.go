package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultUserAgent = "GoTMDBWrapper/1.0.0"
	defaultBaseURL   = "https://api.themoviedb.org/3"
)

var defaultTimeout = 10 * time.Second

// Client manages communication with the TMDb API.
type Client struct {
	// config holds the client configuration.
	config Config

	// httpClient is the underlying HTTP client used for requests.
	httpClient *http.Client

	// userAgent is the user agent string to be used in requests.
	userAgent string

	// baseURL is the parsed base URL for API requests.
	baseURL *url.URL
}

type service struct {
	client *Client
}

func New(config Config) (*Client, error) {
	if config.APIKey == "" {
		return nil, errors.New("tmdb: APIKey is required")
	}

	baseURLStr := config.BaseURL
	if baseURLStr == "" {
		baseURLStr = defaultBaseURL
	}

	parsedBaseURL, err := url.Parse(baseURLStr)
	if err != nil {
		return nil, fmt.Errorf("tmdb: invalid base URL: %w", err)
	}

	// Basic v4 check based on URL prefix (more robust checks might be needed if supporting v4 fully)
	isV4 := strings.HasPrefix(parsedBaseURL.Path, "/4")
	if isV4 && config.BearerToken == "" {
		return nil, errors.New("tmdb: BearerToken is required for v4 API base URL")
	}

	userAgent := config.UserAgent
	httpClient := config.HTTPClient
	if httpClient == nil {
		timeout := config.Timeout
		if timeout == 0 {
			timeout = defaultTimeout
		}

		if userAgent == "" {
			userAgent = defaultUserAgent
		}
		httpClient = &http.Client{Timeout: timeout}
	}

	c := &Client{
		config:     config,
		userAgent:  userAgent,
		httpClient: httpClient,
		baseURL:    parsedBaseURL,
	}

	// TODO: validate API key on creation by making a test call
	// _, err = c.Configuration.GetAPIConfiguration(nil) // Example validation call
	// if err != nil {
	// 	 return nil, fmt.Errorf("tmdb: API key validation failed: %w", err)
	// }

	return c, nil
}

// DoRequest performs the actual HTTP request to the TMDb API.
// If you want to use this method, you should use the client.Client struct. and use the tmdb utils.StructToURLValues to convert the struct parameters to url.Values.
//
//	(I STRONGLY RECOMMEND AGAINST USING THIS METHOD)
//
// Only use this method if you know what you are doing.
// method: HTTP method (e.g., "GET", "POST", "DELETE")
// path: API endpoint path (e.g., "/movie/585511")
// queryParams: URL query parameters (optional)
// requestBody: Data to be sent as the request body for POST/PUT/DELETE (optional)
// responseBody: Pointer to the struct where the successful response should be decoded (optional)
func (c *Client) DoRequest(method, path string, queryParams url.Values, requestBody any, responseBody any) error {
	// Construct the full URL
	relURL, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("tmdb: invalid path %q: %w", path, err)
	}
	fullURL := fmt.Sprintf("%s%s", c.baseURL.String(), relURL.String())

	// Prepares query parameters
	if queryParams == nil {
		queryParams = url.Values{}
	}

	// Adds API key to query parameters if v3
	if !strings.HasPrefix(c.baseURL.Path, "/4") {
		queryParams.Set("api_key", c.config.APIKey)
	}

	if len(queryParams) > 0 {
		encodedParams := queryParams.Encode()
		encodedParams = strings.ReplaceAll(encodedParams, "%2C", ",")
		fullURL += "?" + encodedParams
	}

	// Prepare request body (if any)
	var bodyReader io.Reader
	var reqBytes []byte
	if requestBody != nil {
		reqBytes, err = json.Marshal(requestBody)
		if err != nil {
			return fmt.Errorf("tmdb: failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(reqBytes)
	}

	// Create the HTTP request
	req, err := http.NewRequest(method, fullURL, bodyReader)
	if err != nil {
		return fmt.Errorf("tmdb: failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	if requestBody != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	// Adds Bearer token if configured
	if c.config.BearerToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.config.BearerToken)
	}

	// Execute the request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		// Handles timeout and other context errors
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("tmdb: request timed out: %w", err)
		}
		return fmt.Errorf("tmdb: request failed: %w", err)
	}
	defer resp.Body.Close()

	// Reading the response body
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("tmdb: failed to read response body: %w", err)
	}

	// Checking for API errors (non-2xx status codes)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		apiErr := &TMDBError{}
		// Attempting to decode the TMDB error structure
		if decodeErr := json.Unmarshal(respBytes, apiErr); decodeErr == nil && apiErr.StatusMessage != "" {
			// Using the decoded TMDB error if successful
			apiErr.StatusCode = resp.StatusCode
			return apiErr
		} else {
			// Fallback if decoding fails or message is empty
			return NewTMDBError(resp.StatusCode, fmt.Sprintf("unexpected status code %d with body: %s", resp.StatusCode, string(respBytes)))
		}
	}

	// Decoding successful response body (if expected)
	if responseBody != nil {
		if err := json.Unmarshal(respBytes, responseBody); err != nil {
			return fmt.Errorf("tmdb: failed to decode response body into %T: %w\nBody: %s", responseBody, err, string(respBytes))
		}
	}

	return nil
}
