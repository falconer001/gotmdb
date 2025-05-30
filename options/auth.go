package options

import (
	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/types"
)

type AuthBuilder[T AuthResponse] struct {
	client *client.Client
	path   string
	method string
	body   any
}

type AuthResponse interface {
	*types.StatusResponse |
		*types.SessionResponse |
		*types.RequestTokenResponse |
		*types.GuestSessionResponse
}

func NewAuthBuilder[T AuthResponse](client *client.Client, path string, method string, body any) *AuthBuilder[T] {
	return &AuthBuilder[T]{
		client: client,
		path:   path,
		method: method,
		body:   body,
	}
}

// Exec performs the request and returns the response.
func (n *AuthBuilder[T]) Exec() (*T, error) {
	resp := new(T)

	if n.method == "GET" {
		n.body = nil
	}

	err := n.client.DoRequest(n.method, n.path, nil, n.body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
