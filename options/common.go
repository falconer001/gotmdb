package options

import (
	"fmt"

	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/types"
	"github.com/falconer001/gotmdb/utils"
)

// AppendToResponseBuilder: language + append_to_response (comma-separated)
type AppendToResponseBuilder[T allowedAppendToResponseT] struct {
	client *client.Client
	path   string
	opts   struct {
		Language         *string  `url:"language,omitempty"`
		AppendToResponse []string `url:"append_to_response,omitempty"`
	}
}

type allowedAppendToResponseT interface {
	*types.MovieDetails |
		*types.TVDetails
}

func NewAppendToResponseBuilder[T allowedAppendToResponseT](c *client.Client, path string) *AppendToResponseBuilder[T] {
	return &AppendToResponseBuilder[T]{
		client: c,
		path:   path,
	}
}

// Language sets the language parameter. e.g. "en-US", "fr-FR"
func (b *AppendToResponseBuilder[T]) Language(lang string) *AppendToResponseBuilder[T] {
	b.opts.Language = &lang
	return b
}

// AppendToResponse appends additional data to the response. e.g. "credits", "images", "videos", "keywords", "external_ids".
// Add as many as needed, separated by commas. NOTE: This is a comma-separated string, not an array/slice of strings.
func (b *AppendToResponseBuilder[T]) AppendToResponse(parts ...string) *AppendToResponseBuilder[T] {
	b.opts.AppendToResponse = append(b.opts.AppendToResponse, parts...)
	return b
}

// Exec performs the request and returns the response.
func (b *AppendToResponseBuilder[T]) Exec() (T, error) {
	resp := new(T)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}

	err = b.client.DoRequest("GET", b.path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return *resp, nil
}

// * NO OPTIONS
// noOptsBuilder For endpoints with no query parameters
type NoOptsBuilder[T allowedNoOptsT] struct {
	client *client.Client
	path   string
}

type allowedNoOptsT interface {
	*types.ExternalIDs |
		*types.TVDetails |
		*types.MovieDetails |
		*types.KeywordsResponse |
		*types.ReleaseDatesResponse |
		*types.TranslationsResponse |
		*types.WatchProviderResponse |
		*types.AlternativeTitlesResponse |
		*types.ScreenedTheatricallyResponse
}

func NewNoOptsBuilder[T allowedNoOptsT](c *client.Client, path string) *NoOptsBuilder[T] {
	return &NoOptsBuilder[T]{
		client: c,
		path:   path,
	}
}

// Exec performs the request and returns the response.
func (b *NoOptsBuilder[T]) Exec() (T, error) {
	var zero T
	resp := new(T)

	err := b.client.DoRequest("GET", b.path, nil, nil, resp)
	if err != nil {
		return zero, err
	}

	return *resp, nil
}

// *PAGED BUILDER
// PagedBuilder For endpoints supporting `page`, `language`, and `region`
type PagedBuilder[T allowedPagedT] struct {
	client *client.Client
	path   string
	opts   struct {
		Language *string `url:"language,omitempty"`
		Page     *int    `url:"page,omitempty"`
		Region   *string `url:"region,omitempty"`
		Timezone *string `url:"timezone,omitempty"` // Only for endpoints that support timezone. like (/tv/airing_today)
	}
}

type allowedPagedT interface {
	*types.UpcomingResponse |
		*types.NowPlayingResponse |
		*types.ListPaginatedResults |
		*types.MoviePaginatedResults |
		*types.ReviewPaginatedResults |
		*types.TVShowPaginatedResults
}

func NewPagedBuilder[T allowedPagedT](c *client.Client, path string) *PagedBuilder[T] {
	return &PagedBuilder[T]{
		client: c,
		path:   path,
	}
}

// Language sets the language parameter. e.g. "en-US", "fr-FR"
func (b *PagedBuilder[T]) Language(lang string) *PagedBuilder[T] {
	b.opts.Language = &lang
	return b
}

// Page sets the page parameter. e.g. 1, 2, 3, etc.
func (b *PagedBuilder[T]) Page(p int) *PagedBuilder[T] {
	b.opts.Page = &p
	return b
}

// Region sets the region parameter. e.g. "US", "FR"
// Only for endpoints that support region.
func (b *PagedBuilder[T]) Region(r string) *PagedBuilder[T] {
	b.opts.Region = &r
	return b
}

// Timezone sets the timezone parameter. e.g. "America/New_York", "Europe/Paris"
// Allowed only for endpoints that support timezone. like (tv/airing_today, tv/on_the_air)
func (b *PagedBuilder[T]) Timezone(t string) *PagedBuilder[T] {
	b.opts.Timezone = &t
	return b
}

// Exec performs the request and returns the response.
func (b *PagedBuilder[T]) Exec() (T, error) {
	var zero T
	resp := new(T)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return zero, fmt.Errorf("failed to convert options: %w", err)
	}

	err = b.client.DoRequest("GET", b.path, params, nil, resp)
	if err != nil {
		return zero, err
	}

	return *resp, nil
}

// *LANG BUILDER
// langBuilder For endpoints supporting only `language`
type LangBuilder[T any] struct {
	client *client.Client
	path   string
	opts   struct {
		IncludeImageLanguage *string `url:"include_image_language,omitempty"` // Comma separated ISO 639-1 codes
		Language             *string `url:"language,omitempty"`
	}
}

type allowedLangT interface {
	*types.Credits |
		*types.ImageList |
		*types.VideoList
}

func NewLangBuilder[T any](c *client.Client, path string) *LangBuilder[T] {
	return &LangBuilder[T]{
		client: c,
		path:   path,
	}
}

// Language sets the language parameter. e.g. "en-US", "fr-FR"
func (b *LangBuilder[T]) Language(lang string) *LangBuilder[T] {
	b.opts.Language = &lang
	return b
}

// IncludeImageLanguage sets the include_image_language parameter. e.g. "en-US", "fr-FR"
// Only for /credits, /images, /videos endpoints
func (b *LangBuilder[T]) IncludeImageLanguage(lang string) *LangBuilder[T] {
	b.opts.IncludeImageLanguage = &lang
	return b
}

// Exec performs the request and returns the response.
func (b *LangBuilder[T]) Exec() (T, error) {
	var zero T
	resp := new(T)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return zero, fmt.Errorf("failed to convert options: %w", err)
	}

	err = b.client.DoRequest("GET", b.path, params, nil, resp)
	if err != nil {
		return zero, err
	}

	return *resp, nil
}

// *CHANGES BUILDER
// ChangesBuilder For the /changes endpoint
// Dates must be YYYY-MM-DD
type ChangesBuilder struct {
	client *client.Client
	path   string
	opts   struct {
		StartDate *string `url:"start_date,omitempty"`
		EndDate   *string `url:"end_date,omitempty"`
		Page      *int    `url:"page,omitempty"`
	}
}

func NewChangesBuilder(c *client.Client, path string) *ChangesBuilder {
	return &ChangesBuilder{
		client: c,
		path:   path,
	}
}

// DateRange sets the start and end date parameters.
// Dates must be YYYY-MM-DD
func (b *ChangesBuilder) DateRange(startDate, endDate string) *ChangesBuilder {
	b.opts.StartDate = &startDate
	b.opts.EndDate = &endDate
	return b
}

// Page sets the page parameter.
func (b *ChangesBuilder) Page(p int) *ChangesBuilder {
	b.opts.Page = &p
	return b
}

// Exec performs the request and returns the response.
func (b *ChangesBuilder) Exec() (*types.ItemChangesResponse, error) {
	resp := new(types.ItemChangesResponse)
	params, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options: %w", err)
	}

	err = b.client.DoRequest("GET", b.path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
