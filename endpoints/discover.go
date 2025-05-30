package endpoints

import (
	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/options"
)

type Discover struct {
	Client *client.Client
}

// Discover handles communication with the discover related methods of the TMDb API.
// See: https://developer.themoviedb.org/reference/discover-movie
func (d *Discover) DiscoverMovies() *options.DiscoverMoviesBuilder {
	return options.NewDiscoverMoviesBuilder(d.Client)
}

// DiscoverTV handles communication with the discover tv related methods of the TMDb API.
// See: https://developer.themoviedb.org/reference/discover-tv
func (d *Discover) DiscoverTV() *options.DiscoverTVBuilder {
	return options.NewDiscoverTVBuilder(d.Client)
}
