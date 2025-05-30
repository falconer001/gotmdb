package gotmdb

import (
	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/endpoints"
)

type Config = client.Config

type TMDBClient struct {
	TV       *endpoints.TV
	Search   *endpoints.Search
	Movies   *endpoints.Movies
	Discover *endpoints.Discover
}

func New(config Config) (*TMDBClient, error) {
	var c *client.Client
	var err error
	c, err = client.New(config)
	if err != nil {
		return nil, err
	}

	var tc = &TMDBClient{
		TV:       &endpoints.TV{Client: c},
		Search:   &endpoints.Search{Client: c},
		Movies:   &endpoints.Movies{Client: c},
		Discover: &endpoints.Discover{Client: c},
	}

	return tc, nil
}
