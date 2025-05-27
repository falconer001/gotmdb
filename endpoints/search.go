package endpoints

import (
	"fmt"

	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/types"
	"github.com/falconer001/gotmdb/utils"
)

type SearchService struct {
	Client *client.Client
}

type SearchCompanyOptions struct {
	Query string `url:"query,omitempty"` // Required
	Page  int    `url:"page,omitempty"`
}

// SearchCompanies Search for companies by their original and alternative names.
// I don't really get what this endpoint does. I think it's for networks and production companies. But check the docs if you can understand it.
// See: https://developer.themoviedb.org/reference/search-company
func (ss *SearchService) SearchCompanies(opts *SearchCompanyOptions) (*types.CompanySearchResponse, error) {
	if opts == nil || opts.Query == "" {
		return nil, fmt.Errorf("Query is required for SearchCompanies")
	}
	path := "/search/company"
	resp := new(types.CompanySearchResponse)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ss.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchCollectionOptions represents the available options for the SearchCollections endpoint.
type SearchCollectionOptions struct {
	Query        string `url:"query,omitempty"` // Required
	IncludeAdult bool   `url:"include_adult,omitempty"`
	Language     string `url:"language,omitempty"`
	Page         int    `url:"page,omitempty"`
}

// SearchCollections searches for collections by name.
// See: https://developer.themoviedb.org/reference/search-collection
func (ss *SearchService) SearchCollections(opts *SearchCollectionOptions) (*types.CollectionSearchResponse, error) {
	if opts == nil || opts.Query == "" || len(opts.Query) < 1 {
		return nil, fmt.Errorf("Query is required for SearchCollections")
	}
	path := "/search/collection"
	resp := new(types.CollectionSearchResponse)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ss.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchKeywordOptions represents the available options for the SearchKeywords endpoint.
type SearchKeywordOptions struct {
	Query string `url:"query,omitempty"` // Required
	Page  int    `url:"page,omitempty"`
}

// SearchKeywords searches for keywords by name.
// See: https://developer.themoviedb.org/reference/search-keyword
func (ss *SearchService) SearchKeywords(opts *SearchKeywordOptions) (*types.KeywordSearchResponse, error) {
	if opts == nil || opts.Query == "" {
		return nil, fmt.Errorf("Query is required for SearchKeywords")
	}
	path := "/search/keyword"
	resp := new(types.KeywordSearchResponse)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ss.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchMovieOptions represents the available options for the SearchMovies endpoint.
type SearchMovieOptions struct {
	Query              string `url:"query,omitempty"` // Required
	IncludeAdult       bool   `url:"include_adult,omitempty"`
	Language           string `url:"language,omitempty"`
	PrimaryReleaseYear int    `url:"primary_release_year,omitempty"`
	Page               int    `url:"page,omitempty"`
	Region             string `url:"region,omitempty"`
	Year               int    `url:"year,omitempty"`
}

// SearchMovies searches for movies by title.
// See: https://developer.themoviedb.org/reference/search-movie
func (ss *SearchService) SearchMovies(opts *SearchMovieOptions) (*types.MoviePaginatedResults, error) {
	if opts == nil || opts.Query == "" {
		return nil, fmt.Errorf("Query is required for SearchMovies")
	}
	path := "/search/movie"
	resp := new(types.MoviePaginatedResults)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ss.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchMultiOptions represents the available options for the SearchMulti endpoint.
type SearchMultiOptions struct {
	Query        string `url:"query,omitempty"` // Required
	IncludeAdult bool   `url:"include_adult,omitempty"`
	Language     string `url:"language,omitempty"`
	Page         int    `url:"page,omitempty"`
}

// SearchMulti searches across movies, TV shows, and people in a single request.
// See: https://developer.themoviedb.org/reference/search-multi
func (ss *SearchService) SearchMulti(opts *SearchMultiOptions) (*types.SearchMultiResponse, error) {
	if opts == nil || opts.Query == "" {
		return nil, fmt.Errorf("Query is required for SearchMulti")
	}
	path := "/search/multi"
	resp := new(types.SearchMultiResponse)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ss.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchPersonOptions represents the available options for the SearchPeople endpoint.
type SearchPersonOptions struct {
	Query        string `url:"query,omitempty"` // Required
	IncludeAdult bool   `url:"include_adult,omitempty"`
	Language     string `url:"language,omitempty"`
	Page         int    `url:"page,omitempty"`
}

// SearchPeople searches for people by name.
// See: https://developer.themoviedb.org/reference/search-person
func (ss *SearchService) SearchPeople(opts *SearchPersonOptions) (*types.PersonPaginatedResults, error) {
	if opts == nil || opts.Query == "" {
		return nil, fmt.Errorf("Query is required for SearchPeople")
	}
	path := "/search/person"
	resp := new(types.PersonPaginatedResults)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ss.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchTVOptions represents the available options for the SearchTVShows endpoint.
type SearchTVOptions struct {
	Query            string `url:"query,omitempty"` // Required
	IncludeAdult     bool   `url:"include_adult,omitempty"`
	Language         string `url:"language,omitempty"`
	FirstAirDateYear int    `url:"first_air_date_year,omitempty"`
	Page             int    `url:"page,omitempty"`
	Year             int    `url:"year,omitempty"` // Deprecated, use FirstAirDateYear
}

// SearchTVShows searches for TV shows by title.
// See: https://developer.themoviedb.org/reference/search-tv
func (ss *SearchService) SearchTVShows(opts *SearchTVOptions) (*types.TVShowPaginatedResults, error) {
	if opts == nil || opts.Query == "" {
		return nil, fmt.Errorf("Query is required for SearchTVShows")
	}
	path := "/search/tv"
	resp := new(types.TVShowPaginatedResults)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ss.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
