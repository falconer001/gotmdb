package options

import (
	"fmt"
	"log"
	"slices"

	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/types"
	"github.com/falconer001/gotmdb/utils"
)

// --- Search Movies --- //

type searchMoviesOptions struct {
	Query              string  `url:"query"` // Required, but handled in path/constructor
	IncludeAdult       *bool   `url:"include_adult,omitempty"`
	Language           *string `url:"language,omitempty"`
	PrimaryReleaseYear *int    `url:"primary_release_year,omitempty"`
	Page               *int    `url:"page,omitempty"`
	Region             *string `url:"region,omitempty"`
	Year               *int    `url:"year,omitempty"`
}

type SearchMoviesBuilder struct {
	client *client.Client
	opts   searchMoviesOptions
}

func NewSearchMoviesBuilder(c *client.Client, query string) *SearchMoviesBuilder {
	return &SearchMoviesBuilder{
		client: c,
		opts:   searchMoviesOptions{Query: query},
	}
}

func (b *SearchMoviesBuilder) IncludeAdult(include bool) *SearchMoviesBuilder {
	b.opts.IncludeAdult = &include
	return b
}

func (b *SearchMoviesBuilder) Language(lang string) *SearchMoviesBuilder {
	b.opts.Language = &lang
	return b
}

func (b *SearchMoviesBuilder) PrimaryReleaseYear(year int) *SearchMoviesBuilder {
	b.opts.PrimaryReleaseYear = &year
	return b
}

func (b *SearchMoviesBuilder) Page(page int) *SearchMoviesBuilder {
	b.opts.Page = &page
	return b
}

func (b *SearchMoviesBuilder) Region(region string) *SearchMoviesBuilder {
	b.opts.Region = &region
	return b
}

func (b *SearchMoviesBuilder) Year(year int) *SearchMoviesBuilder {
	b.opts.Year = &year
	return b
}

func (b *SearchMoviesBuilder) Exec() (*types.MoviePaginatedResults, error) {
	path := "/search/movie"
	resp := new(types.MoviePaginatedResults)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}
	err = b.client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// --- Search TV --- //

type searchTVOptions struct {
	Query            string  `url:"query"`
	IncludeAdult     *bool   `url:"include_adult,omitempty"`
	Language         *string `url:"language,omitempty"`
	FirstAirDateYear *int    `url:"first_air_date_year,omitempty"`
	Page             *int    `url:"page,omitempty"`
	Year             *int    `url:"year,omitempty"` // Deprecated, use FirstAirDateYear
}

type SearchTVBuilder struct {
	client *client.Client
	opts   searchTVOptions
}

func NewSearchTVBuilder(c *client.Client, query string) *SearchTVBuilder {
	return &SearchTVBuilder{
		client: c,
		opts:   searchTVOptions{Query: query},
	}
}

func (b *SearchTVBuilder) IncludeAdult(include bool) *SearchTVBuilder {
	b.opts.IncludeAdult = &include
	return b
}

func (b *SearchTVBuilder) Language(lang string) *SearchTVBuilder {
	b.opts.Language = &lang
	return b
}

func (b *SearchTVBuilder) FirstAirDateYear(year int) *SearchTVBuilder {
	b.opts.FirstAirDateYear = &year
	return b
}

func (b *SearchTVBuilder) Page(page int) *SearchTVBuilder {
	b.opts.Page = &page
	return b
}

// Year sets the deprecated year parameter. Use FirstAirDateYear instead.
func (b *SearchTVBuilder) Year(year int) *SearchTVBuilder {
	b.opts.Year = &year
	return b
}

func (b *SearchTVBuilder) Exec() (*types.TVShowPaginatedResults, error) {
	path := "/search/tv"
	resp := new(types.TVShowPaginatedResults)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}
	err = b.client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// --- Search Multi --- //

type searchMultiOptions struct {
	Query         string  `url:"query"`
	IncludePeople *bool   `url:"include_people,omitempty"` //default is false
	IncludeAdult  *bool   `url:"include_adult,omitempty"`
	Language      *string `url:"language,omitempty"`
	Page          *int    `url:"page,omitempty"`
}

type SearchMultiBuilder struct {
	client *client.Client
	opts   searchMultiOptions
}

func NewSearchMultiBuilder(c *client.Client, query string) *SearchMultiBuilder {
	return &SearchMultiBuilder{
		client: c,
		opts:   searchMultiOptions{Query: query},
	}
}

func (b *SearchMultiBuilder) IncludePeople(include bool) *SearchMultiBuilder {
	b.opts.IncludePeople = &include
	return b
}

func (b *SearchMultiBuilder) IncludeAdult(include bool) *SearchMultiBuilder {
	b.opts.IncludeAdult = &include
	return b
}

func (b *SearchMultiBuilder) Language(lang string) *SearchMultiBuilder {
	b.opts.Language = &lang
	return b
}

func (b *SearchMultiBuilder) Page(page int) *SearchMultiBuilder {
	b.opts.Page = &page
	return b
}

// Exec performs the search multi request and returns the response.
// If IncludePeople is false, person results are removed from the response (this is the default mode).
func (b *SearchMultiBuilder) Exec() (*types.SearchMultiResponse, error) {
	path := "/search/multi"
	resp := new(types.SearchMultiResponse)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}
	err = b.client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}

	if b.opts.IncludePeople == nil || !*b.opts.IncludePeople {
		var totalRemoved int
		for i := len(resp.Results) - 1; i >= 0; i-- {
			if resp.Results[i].MediaType == "person" {
				resp.Results = slices.Delete(resp.Results, i, i+1)
				totalRemoved++
			}
		}
		log.Printf("tmdb: Removed %d person results\n", totalRemoved)
	}

	return resp, nil
}

// --- Search Companies --- //

type searchCompaniesOptions struct {
	Query string `url:"query"`
	Page  *int   `url:"page,omitempty"`
}

type SearchCompaniesBuilder struct {
	client *client.Client
	opts   searchCompaniesOptions
}

func NewSearchCompaniesBuilder(c *client.Client, query string) *SearchCompaniesBuilder {
	return &SearchCompaniesBuilder{
		client: c,
		opts:   searchCompaniesOptions{Query: query},
	}
}

func (b *SearchCompaniesBuilder) Page(page int) *SearchCompaniesBuilder {
	b.opts.Page = &page
	return b
}

func (b *SearchCompaniesBuilder) Exec() (*types.CompanySearchResponse, error) {
	path := "/search/company"
	resp := new(types.CompanySearchResponse)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}
	err = b.client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// --- Search Collections --- //

type searchCollectionsOptions struct {
	Query        string  `url:"query"`
	IncludeAdult *bool   `url:"include_adult,omitempty"`
	Language     *string `url:"language,omitempty"`
	Page         *int    `url:"page,omitempty"`
}

type SearchCollectionsBuilder struct {
	client *client.Client
	opts   searchCollectionsOptions
}

func NewSearchCollectionsBuilder(c *client.Client, query string) *SearchCollectionsBuilder {
	return &SearchCollectionsBuilder{
		client: c,
		opts:   searchCollectionsOptions{Query: query},
	}
}

func (b *SearchCollectionsBuilder) IncludeAdult(include bool) *SearchCollectionsBuilder {
	b.opts.IncludeAdult = &include
	return b
}

func (b *SearchCollectionsBuilder) Language(lang string) *SearchCollectionsBuilder {
	b.opts.Language = &lang
	return b
}

func (b *SearchCollectionsBuilder) Page(page int) *SearchCollectionsBuilder {
	b.opts.Page = &page
	return b
}

func (b *SearchCollectionsBuilder) Exec() (*types.CollectionSearchResponse, error) {
	path := "/search/collection"
	resp := new(types.CollectionSearchResponse)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}
	err = b.client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// --- Search Keywords --- //

type searchKeywordsOptions struct {
	Query string `url:"query"`
	Page  *int   `url:"page,omitempty"`
}

type SearchKeywordsBuilder struct {
	client *client.Client
	opts   searchKeywordsOptions
}

func NewSearchKeywordsBuilder(c *client.Client, query string) *SearchKeywordsBuilder {
	return &SearchKeywordsBuilder{
		client: c,
		opts:   searchKeywordsOptions{Query: query},
	}
}

func (b *SearchKeywordsBuilder) Page(page int) *SearchKeywordsBuilder {
	b.opts.Page = &page
	return b
}

func (b *SearchKeywordsBuilder) Exec() (*types.KeywordSearchResponse, error) {
	path := "/search/keyword"
	resp := new(types.KeywordSearchResponse)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}
	err = b.client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// --- Search People --- //

type searchPeopleOptions struct {
	Query        string  `url:"query"`
	IncludeAdult *bool   `url:"include_adult,omitempty"`
	Language     *string `url:"language,omitempty"`
	Page         *int    `url:"page,omitempty"`
}

type SearchPeopleBuilder struct {
	client *client.Client
	opts   searchPeopleOptions
}

func NewSearchPeopleBuilder(c *client.Client, query string) *SearchPeopleBuilder {
	return &SearchPeopleBuilder{
		client: c,
		opts:   searchPeopleOptions{Query: query},
	}
}

func (b *SearchPeopleBuilder) IncludeAdult(include bool) *SearchPeopleBuilder {
	b.opts.IncludeAdult = &include
	return b
}

func (b *SearchPeopleBuilder) Language(lang string) *SearchPeopleBuilder {
	b.opts.Language = &lang
	return b
}

func (b *SearchPeopleBuilder) Page(page int) *SearchPeopleBuilder {
	b.opts.Page = &page
	return b
}

func (b *SearchPeopleBuilder) Exec() (*types.PersonPaginatedResults, error) {
	path := "/search/person"
	resp := new(types.PersonPaginatedResults)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}
	err = b.client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
