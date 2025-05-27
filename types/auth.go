package types

// GuestSessionResponse represents the response when creating a guest session.
// See: https://developer.themoviedb.org/reference/authentication-create-guest-session
type GuestSessionResponse struct {
	Success        bool   `json:"success"`
	GuestSessionID string `json:"guest_session_id"`
	ExpiresAt      string `json:"expires_at"` // Timestamp, e.g., "2024-01-01 12:00:00 UTC"
}

// RequestTokenResponse represents the response when creating a request token.
// See: https://developer.themoviedb.org/reference/authentication-create-request-token
type RequestTokenResponse struct {
	Success       bool   `json:"success"`
	ExpiresAt     string `json:"expires_at"` // Timestamp
	RequestToken string `json:"request_token"`
}

// CreateSessionRequest is the request body for creating a session ID from a request token.
// See: https://developer.themoviedb.org/reference/authentication-create-session
type CreateSessionRequest struct {
	RequestToken string `json:"request_token"`
}

// SessionResponse represents the response when creating a session ID.
// See: https://developer.themoviedb.org/reference/authentication-create-session
type SessionResponse struct {
	Success   bool   `json:"success"`
	SessionID string `json:"session_id"`
}

// ValidateWithLoginRequest is the request body for the deprecated username/password login flow.
// Note: Use of this endpoint is discouraged; prefer the browser-based authentication flow.
// See: https://developer.themoviedb.org/reference/authentication-validate-user
type ValidateWithLoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	RequestToken string `json:"request_token"`
}

// ConvertV4TokenRequest is the request body for creating a v3 session ID from a v4 access token.
// See: https://developer.themoviedb.org/reference/authentication-create-session-from-v4-token
type ConvertV4TokenRequest struct {
	AccessToken string `json:"access_token"` // v4 access token
}

// DeleteSessionRequest is the request body for deleting a session ID.
// See: https://developer.themoviedb.org/reference/authentication-delete-session
type DeleteSessionRequest struct {
	SessionID string `json:"session_id"`
}

