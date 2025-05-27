package client

import (
	"fmt"
)

// TMDBError represents an error response from the TMDB API.
// It includes the status code and message provided by the API.
// See: https://developer.themoviedb.org/docs/errors
type TMDBError struct {
	StatusCode    int    `json:"status_code"`    // The HTTP status code.
	StatusMessage string `json:"status_message"` // The error message from TMDB.
	// Success field is sometimes present in error responses, often false.
	// I primarily rely on HTTP status code for error detection.
	// Success       *bool  `json:"success,omitempty"`
}

// Error returns the string representation of the TMDBError.
func (e *TMDBError) Error() string {
	return fmt.Sprintf("tmdb: API error (HTTP %d): %s", e.StatusCode, e.StatusMessage)
}

// Is checks if the target error is a TMDBError with the same status code.
// This helps with error checking using errors.Is.
func (e *TMDBError) Is(target error) bool {
	tmdbErr, ok := target.(*TMDBError)
	if !ok {
		return false
	}
	return e.StatusCode == tmdbErr.StatusCode
}

// NewTMDBError creates a new TMDBError.
func NewTMDBError(code int, message string) *TMDBError {
	return &TMDBError{
		StatusCode:    code,
		StatusMessage: message,
	}
}
