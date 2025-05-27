package client

import (
	"net/http"
	"time"
)

// Config holds the configuration options for creating a new TMDb API client.
// All fields are optional except where noted.
type Config struct {
	// APIKey is your TMDb API v3 key (required for v3 API).
	// Get one from: https://www.themoviedb.org/settings/api
	APIKey string

	// BearerToken is your TMDb API Read Access Token (required for profile and auth related endpoints).
	// It gives easy access throughout the API instaed of generating a session id everytime its needed.
	// Generate one using your v3 API key at: https://www.themoviedb.org/settings/api
	BearerToken string

	// BaseURL is the base URL for TMDb API requests.
	// Defaults to "https://api.themoviedb.org/3" for v3 API.
	// For v4 API, use "https://api.themoviedb.org/4" and ensure BearerToken is set.
	BaseURL string

	// HTTPClient allows specifying a custom *http.Client.
	// If nil, a default client with a 10-second timeout will be used.
	HTTPClient *http.Client

	// Timeout specifies a time limit for requests made by the HTTP client.
	// Only used when HTTPClient is nil. Default is 10 seconds.
	Timeout time.Duration

	// UserAgent is the user agent string to be used in requests.
	// Defaults to "GoTMDBWrapper/{version}"
	UserAgent string

	// UseProxy enables proxy support when set to true.
	// Note: Proxy support is not yet implemented in this design.
	// UseProxy bool

	// Proxies contains a list of proxy URLs to be used when UseProxy is true.
	// Format: []string{"http://proxy1:port", "http://proxy2:port"}
	// Note: Proxy support is not yet implemented in this design.
	// Proxies []string
}
