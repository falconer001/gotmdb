package endpoints

import (
	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/options"
	"github.com/falconer001/gotmdb/types"
)

type Auth struct {
	Client *client.Client
}

// CreateRequestToken creates a temporary request token that can be used to validate
// a TMDb user login. More details about how this works can be found in the
// v3/authentication/how-do-i-generate-a-session-id documentation.
// Make sure to prompt the user to approve the request token in their browser.
// See: https://developer.themoviedb.org/reference/authentication-create-request-token
func (a *Auth) CreateRequestToken() *options.AuthBuilder[*types.RequestTokenResponse] {
	return options.NewAuthBuilder[*types.RequestTokenResponse](a.Client, "/authentication/token/new", "GET", nil)
}

// CreateGuestSession creates a temporary guest session.
// Guest sessions are useful for users who don't have or don't want to
// sign up for a TMDb account but still want to rate movies and TV shows.
// Guest sessions expire after 24 hours.
// See: https://developer.themoviedb.org/reference/authentication-create-guest-session
func (a *Auth) CreateGuestSession() *options.AuthBuilder[*types.GuestSessionResponse] {
	return options.NewAuthBuilder[*types.GuestSessionResponse](a.Client, "/authentication/guest_session/new", "GET", nil)
}

// CreateSession creates a session ID for user based authentication.
// A request token is required to generate the session ID.
// See: https://developer.themoviedb.org/reference/authentication-create-session
func (a *Auth) CreateSession(body types.CreateSessionRequest) *options.AuthBuilder[*types.SessionResponse] {
	return options.NewAuthBuilder[*types.SessionResponse](a.Client, "/authentication/session/new", "POST", body)
}

// ValidateWithLogin validates a request token by authenticating with a TMDb username and password.
// Note: This method is deprecated and is not the preferred way to obtain user authentication.
// Use the browser approval method instead.
// See: https://developer.themoviedb.org/reference/authentication-validate-user
func (a *Auth) ValidateWithLogin(body types.ValidateWithLoginRequest) *options.AuthBuilder[*types.RequestTokenResponse] {
	return options.NewAuthBuilder[*types.RequestTokenResponse](a.Client, "/authentication/token/validate_with_login", "POST", body)
}

// CreateSessionFromV4 creates a v3 session ID by validating a v4 access token.
// See: https://developer.themoviedb.org/reference/authentication-create-session-from-v4-token
func (a *Auth) CreateSessionFromV4(body types.ConvertV4TokenRequest) *options.AuthBuilder[*types.SessionResponse] {
	return options.NewAuthBuilder[*types.SessionResponse](a.Client, "/authentication/session/convert/4", "POST", body)
}

// DeleteSession logs out of a session.
// See: https://developer.themoviedb.org/reference/authentication-delete-session
func (a *Auth) DeleteSession(body types.DeleteSessionRequest) *options.AuthBuilder[*types.StatusResponse] {
	return options.NewAuthBuilder[*types.StatusResponse](a.Client, "/authentication/session", "DELETE", body)
}
