package gotmdb

import (
	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/endpoints"
)

// Config holds the configuration options for creating a new TMDb API client.
type Config = client.Config

type Client struct {
	*client.Client

	AuthService   *endpoints.AuthService
	MoviesService *endpoints.MoviesService
	SearchService *endpoints.SearchService
}

func New(config Config) (*Client, error) {
	c, err := client.New(config)
	if err != nil {
		return nil, err
	}

	tc := &Client{
		Client: c,

		AuthService:   &endpoints.AuthService{Client: c},
		MoviesService: &endpoints.MoviesService{Client: c},
		SearchService: &endpoints.SearchService{Client: c},
	}

	return tc, nil
}
