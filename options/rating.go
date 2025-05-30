package options

import (
	"fmt"

	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/types"
	"github.com/falconer001/gotmdb/utils"
)

type StateSessionBuilder[T any] struct {
	client *client.Client
	path   string
	method string
	body   any
	opts   struct {
		ForGuest       bool    //if false uses set client bearer token, if true uses sessionId/guestSessionId
		SessionID      *string `url:"session_id,omitempty"`
		GuestSessionID *string `url:"guest_session_id,omitempty"`
	}
}

type allowedStateSessionT interface {
	*types.StatusResponse | *types.AccountState
}

func NewStateSessionBuilder[T allowedStateSessionT](c *client.Client, path string, method string, body any) *StateSessionBuilder[T] {
	return &StateSessionBuilder[T]{
		client: c,
		body:   body,
		path:   path,
		method: method,
	}
}

// ForGuest sets the forGuest parameter.
// This a bool that determines whether to use the client's bearer token or a session ID.
func (b *StateSessionBuilder[T]) ForGuest() *StateSessionBuilder[T] {
	b.opts.ForGuest = true
	return b
}

// SessionID sets the session ID parameter.
// See: https://developer.themoviedb.org/reference/movie-add-rating
func (b *StateSessionBuilder[T]) SessionID(id string) *StateSessionBuilder[T] {
	b.opts.SessionID = &id
	return b
}

// GuestSessionID sets the guest session ID parameter.
func (b *StateSessionBuilder[T]) GuestSessionID(id string) *StateSessionBuilder[T] {
	b.opts.GuestSessionID = &id
	return b
}

// Exec performs the request and returns the response.
func (b *StateSessionBuilder[T]) Exec() (T, error) {
	var zero T

	if b.opts.ForGuest {
		if b.opts.SessionID == nil || b.opts.GuestSessionID == nil {
			return zero, fmt.Errorf("either SessionID or GuestSessionID is required for rating")
		}
	}

	resp := new(T)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return zero, fmt.Errorf("failed to convert options: %w", err)
	}

	err = b.client.DoRequest(b.method, b.path, params, b.body, resp)
	if err != nil {
		return zero, err
	}
	return *resp, nil
}
