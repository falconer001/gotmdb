package endpoints

import (
	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/types"
)

// AuthService handles communication with the authentication related
// methods of the TMDb API.
// See: https://developer.themoviedb.org/reference/authentication
type AuthService struct {
	Client *client.Client
}

// CreateGuestSession creates a temporary guest session.
// Guest sessions are useful for users who don't have or don't want to
// sign up for a TMDb account but still want to rate movies and TV shows.
// Guest sessions expire after 24 hours.
// See: https://developer.themoviedb.org/reference/authentication-create-guest-session
func (as *AuthService) CreateGuestSession() (*types.GuestSessionResponse, error) {
	path := "/authentication/guest_session/new"
	resp := new(types.GuestSessionResponse)
	err := as.Client.DoRequest("GET", path, nil, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateRequestToken creates a temporary request token that can be used to validate
// a TMDb user login. More details about how this works can be found in the
// v3/authentication/how-do-i-generate-a-session-id documentation.
// Make sure to prompt the user to approve the request token in their browser.
// See: https://developer.themoviedb.org/reference/authentication-create-request-token
func (as *AuthService) CreateRequestToken() (*types.RequestTokenResponse, error) {
	path := "/authentication/token/new"
	resp := new(types.RequestTokenResponse)
	err := as.Client.DoRequest("GET", path, nil, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateSession creates a session ID for user based authentication.
// A request token is required to generate the session ID.
// See: https://developer.themoviedb.org/reference/authentication-create-session
func (as *AuthService) CreateSession(body types.CreateSessionRequest) (*types.SessionResponse, error) {
	path := "/authentication/session/new"
	resp := new(types.SessionResponse)
	err := as.Client.DoRequest("POST", path, nil, body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ValidateWithLogin validates a request token by authenticating with a TMDb username and password.
// Note: This method is deprecated and is not the preferred way to obtain user authentication.
// Use the browser-based approval flow instead.
// See: https://developer.themoviedb.org/reference/authentication-validate-user
func (as *AuthService) ValidateWithLogin(body types.ValidateWithLoginRequest) (*types.RequestTokenResponse, error) {
	path := "/authentication/token/validate_with_login"
	resp := new(types.RequestTokenResponse) // Response is the same as CreateRequestToken
	err := as.Client.DoRequest("POST", path, nil, body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateSessionFromV4 creates a v3 session ID by validating a v4 access token.
// See: https://developer.themoviedb.org/reference/authentication-create-session-from-v4-token
func (as *AuthService) CreateSessionFromV4(body types.ConvertV4TokenRequest) (*types.SessionResponse, error) {
	path := "/authentication/session/convert/4"
	resp := new(types.SessionResponse)
	err := as.Client.DoRequest("POST", path, nil, body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteSession logs out of a session.
// See: https://developer.themoviedb.org/reference/authentication-delete-session
func (as *AuthService) DeleteSession(body types.DeleteSessionRequest) (*types.StatusResponse, error) {
	path := "/authentication/session"
	resp := new(types.StatusResponse)
	err := as.Client.DoRequest("DELETE", path, nil, body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
