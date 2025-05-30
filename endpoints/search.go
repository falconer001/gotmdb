package endpoints

import (
	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/options"
)

// Search handles communication with the search related methods of the TMDb API.
type Search struct {
	*client.Client
}

// Movies initiates the request builder for searching movies.
// Required parameter: query.
// See: https://developer.themoviedb.org/reference/search-movie
func (s *Search) Movies(query string) *options.SearchMoviesBuilder {
	return options.NewSearchMoviesBuilder(s.Client, query)
}

// TV initiates the request builder for searching TV shows.
// Required parameter: query.
// See: https://developer.themoviedb.org/reference/search-tv
func (s *Search) TV(query string) *options.SearchTVBuilder {
	return options.NewSearchTVBuilder(s.Client, query)
}

// Multi initiates the request builder for searching movies, TV shows, and people.
// Default mode removes person results from the response.
// To include person results, use the IncludePeople option.
// Required parameter: query.
// See: https://developer.themoviedb.org/reference/search-multi
func (s *Search) Multi(query string) *options.SearchMultiBuilder {
	return options.NewSearchMultiBuilder(s.Client, query)
}

// Companies initiates the request builder for searching companies.
// Required parameter: query.
// See: https://developer.themoviedb.org/reference/search-company
func (s *Search) Companies(query string) *options.SearchCompaniesBuilder {
	return options.NewSearchCompaniesBuilder(s.Client, query)
}

// Collections initiates the request builder for searching collections.
// Required parameter: query.
// See: https://developer.themoviedb.org/reference/search-collection
func (s *Search) Collections(query string) *options.SearchCollectionsBuilder {
	return options.NewSearchCollectionsBuilder(s.Client, query)
}

// Keywords initiates the request builder for searching keywords.
// Required parameter: query.
// See: https://developer.themoviedb.org/reference/search-keyword
func (s *Search) Keywords(query string) *options.SearchKeywordsBuilder {
	return options.NewSearchKeywordsBuilder(s.Client, query)
}

// People initiates the request builder for searching people.
// Required parameter: query.
// See: https://developer.themoviedb.org/reference/search-person
func (s *Search) People(query string) *options.SearchPeopleBuilder {
	return options.NewSearchPeopleBuilder(s.Client, query)
}
