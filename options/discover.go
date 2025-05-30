package options

import (
	"fmt"
	"net/url"
	"time"

	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/types"
	"github.com/falconer001/gotmdb/utils"
)

type allowedBaseBuilderTypes interface {
	isBuilder()
}

type BaseOpts struct {
	Language                   *string  `url:"language,omitempty"`
	Region                     *string  `url:"region,omitempty"`
	Page                       *int     `url:"page,omitempty"`
	SortBy                     *string  `url:"sort_by,omitempty"`
	WithOriginCountry          *string  `url:"with_origin_country,omitempty"`
	WithOriginalLanguage       *string  `url:"with_original_language,omitempty"`
	WithGenres                 []string `url:"with_genres,omitempty,comma"`
	WithoutGenres              []string `url:"without_genres,omitempty,comma"`
	WithKeywords               []string `url:"with_keywords,omitempty,pipe"`
	WithoutKeywords            []string `url:"without_keywords,omitempty,pipe"`
	WithWatchProviders         []string `url:"with_watch_providers,omitempty,pipe"`
	WithCompanies              []string `url:"with_companies,omitempty,pipe"`
	WithoutWatchProviders      []string `url:"without_watch_providers,omitempty,pipe"`
	WatchRegion                *string  `url:"watch_region,omitempty"`
	WithWatchMonetizationTypes []string `url:"with_watch_monetization_types,omitempty"`
	Timezone                   *string  `url:"timezone,omitempty"`
	IncludeAdult               *bool    `url:"include_adult,omitempty"`
}

type BaseDiscoverBuilder[T allowedBaseBuilderTypes] struct {
	BaseOpts
	self T `url:"-"`
}

func (b *BaseDiscoverBuilder[T]) isBuilder() {
	// empty method to satisfy the interface
}

// Language sets the language parameter. e.g. "en-US", "fr-FR" (default: en-US)
func (b *BaseDiscoverBuilder[T]) Language(lang string) T {
	b.BaseOpts.Language = &lang
	return b.self
}

func (b *BaseDiscoverBuilder[T]) Region(region string) T {
	b.BaseOpts.Region = &region
	return b.self
}

func (b *BaseDiscoverBuilder[T]) Page(page int) T {
	b.BaseOpts.Page = &page
	return b.self
}

func (b *BaseDiscoverBuilder[T]) WithCompanies(ids ...string) T {
	b.BaseOpts.WithCompanies = append(b.BaseOpts.WithCompanies, ids...)
	return b.self
}

// SortBy sets the sort_by parameter. e.g. "original_title.asc", "original_title.desc", "popularity.asc", "popularity.desc", "revenue.asc", "revenue.desc", "primary_release_date.asc", "title.asc", "title.desc", "primary_release_date.desc", "vote_average.asc", "vote_average.desc", "vote_count.asc", "vote_count.desc" Defaults to popularity.desc
func (b *BaseDiscoverBuilder[T]) SortBy(sort string) T {
	b.BaseOpts.SortBy = &sort
	return b.self
}

// WithOriginCountry adds a country to the withOriginCountry field.
func (b *BaseDiscoverBuilder[T]) WithOriginCountry(country string) T {
	b.BaseOpts.WithOriginCountry = &country
	return b.self
}

// WithOriginalLanguage adds a language to the withOriginalLanguage field.
func (b *BaseDiscoverBuilder[T]) WithOriginalLanguage(lang string) T {
	b.BaseOpts.WithOriginalLanguage = &lang
	return b.self
}

// WithGenres adds genres to the withGenres slice. Or not yet implemented.
func (b *BaseDiscoverBuilder[T]) WithGenres(genres ...string) T {
	b.BaseOpts.WithGenres = append(b.BaseOpts.WithGenres, genres...)
	return b.self
}

func (b *BaseDiscoverBuilder[T]) WithoutGenres(genres ...string) T {
	b.BaseOpts.WithoutGenres = append(b.BaseOpts.WithoutGenres, genres...)
	return b.self
}

func (b *BaseDiscoverBuilder[T]) WithKeywords(keys ...string) T {
	b.BaseOpts.WithKeywords = append(b.BaseOpts.WithKeywords, keys...)
	return b.self
}

func (b *BaseDiscoverBuilder[T]) WithoutKeywords(keys ...string) T {
	b.BaseOpts.WithoutKeywords = append(b.BaseOpts.WithoutKeywords, keys...)
	return b.self
}

// use in conjunction with watch_region
func (b *BaseDiscoverBuilder[T]) WithWatchProviders(ids ...string) T {
	b.BaseOpts.WithWatchProviders = append(b.BaseOpts.WithWatchProviders, ids...)
	return b.self
}

func (b *BaseDiscoverBuilder[T]) WithoutWatchProviders(ids ...string) T {
	b.BaseOpts.WithoutWatchProviders = append(b.BaseOpts.WithoutWatchProviders, ids...)
	return b.self
}

// use in conjunction with with_watch_monetization_types or with_watch_providers
func (b *BaseDiscoverBuilder[T]) WatchRegion(region string) T {
	b.BaseOpts.WatchRegion = &region
	return b.self
}

// possible values are: [flatrate, free, ads, rent, buy] use in conjunction with watch_region. Comma seperated string
func (b *BaseDiscoverBuilder[T]) WithWatchMonetizationTypes(m ...string) T {
	b.BaseOpts.WithWatchMonetizationTypes = append(b.BaseOpts.WithWatchMonetizationTypes, m...)
	return b.self
}

// Timezone sets the timezone parameter. e.g. "America/New_York"
func (b *BaseDiscoverBuilder[T]) Timezone(tz string) T {
	b.BaseOpts.Timezone = &tz
	return b.self
}

// IncludeAdult sets the include_adult parameter. e.g. true, false (default: false)
func (b *BaseDiscoverBuilder[T]) IncludeAdult(v bool) T {
	b.BaseOpts.IncludeAdult = &v
	return b.self
}

// Builder for /discover/movie
type DiscoverMoviesBuilder struct {
	client *client.Client `json:"-" url:"-"`
	*BaseDiscoverBuilder[*DiscoverMoviesBuilder]
	opts struct {
		Year                  *int     `url:"year,omitempty"`
		ReleaseDateGTE        *string  `url:"release_date.gte,omitempty"`
		ReleaseDateLTE        *string  `url:"release_date.lte,omitempty"`
		PrimaryReleaseYear    *int     `url:"primary_release_year,omitempty"`
		PrimaryReleaseDateGTE *string  `url:"primary_release_date.gte,omitempty"`
		PrimaryReleaseDateLTE *string  `url:"primary_release_date.lte,omitempty"`
		Certification         *string  `url:"certification,omitempty"`
		CertificationGTE      *string  `url:"certification.gte,omitempty"`
		CertificationLTE      *string  `url:"certification.lte,omitempty"`
		CertificationCountry  *string  `url:"certification_country,omitempty"`
		IncludeVideo          *bool    `url:"include_video,omitempty"`
		WithoutCompanies      []string `url:"without_companies,omitempty,pipe"`
		WithCast              []string `url:"with_cast,omitempty,pipe"`
		WithCrew              []string `url:"with_crew,omitempty,pipe"`
		WithPeople            []string `url:"with_people,omitempty,pipe"`
		WithReleaseType       *int     `url:"with_release_type,omitempty"`
		VoteAverageGTE        *float64 `url:"vote_average.gte,omitempty"`
		VoteAverageLTE        *float64 `url:"vote_average.lte,omitempty"`
		VoteCountGTE          *int     `url:"vote_count.gte,omitempty"`
		VoteCountLTE          *int     `url:"vote_count.lte,omitempty"`
		WithRuntimeGTE        *int     `url:"with_runtime.gte,omitempty"`
		WithRuntimeLTE        *int     `url:"with_runtime.lte,omitempty"`
	}
}

func (b *DiscoverMoviesBuilder) isBuilder() {}

// NewDiscoverMoviesBuilder initializes the movie discover builder.
func NewDiscoverMoviesBuilder(c *client.Client) *DiscoverMoviesBuilder {
	b := &DiscoverMoviesBuilder{}
	b.client = c
	b.BaseDiscoverBuilder = &BaseDiscoverBuilder[*DiscoverMoviesBuilder]{
		self: b,
	}
	return b
}

func (b *DiscoverMoviesBuilder) Year(y int) *DiscoverMoviesBuilder { b.opts.Year = &y; return b }

func (b *DiscoverMoviesBuilder) PrimaryReleaseYear(y int) *DiscoverMoviesBuilder {
	b.opts.PrimaryReleaseYear = &y
	return b
}

func (b *DiscoverMoviesBuilder) PrimaryReleaseDateGTE(t time.Time) *DiscoverMoviesBuilder {
	s := t.Format("2006-01-02")
	b.opts.PrimaryReleaseDateGTE = &s
	return b
}

func (b *DiscoverMoviesBuilder) PrimaryReleaseDateLTE(t time.Time) *DiscoverMoviesBuilder {
	s := t.Format("2006-01-02")
	b.opts.PrimaryReleaseDateLTE = &s
	return b
}

func (b *DiscoverMoviesBuilder) ReleaseDateGTE(t time.Time) *DiscoverMoviesBuilder {
	s := t.Format("2006-01-02")
	b.opts.ReleaseDateGTE = &s
	return b
}

func (b *DiscoverMoviesBuilder) ReleaseDateLTE(t time.Time) *DiscoverMoviesBuilder {
	s := t.Format("2006-01-02")
	b.opts.ReleaseDateLTE = &s
	return b
}

// Certification sets the certification parameter. e.g. "R", "PG-13"
// use in conjunction with region
func (b *DiscoverMoviesBuilder) Certification(c string) *DiscoverMoviesBuilder {
	b.opts.Certification = &c
	return b
}

// CertificationGTE sets the certification.gte parameter. e.g. "R", "PG-13"
// use in conjunction with region
func (b *DiscoverMoviesBuilder) CertificationGTE(c string) *DiscoverMoviesBuilder {
	b.opts.CertificationGTE = &c
	return b
}

// CertificationLTE sets the certification.lte parameter. e.g. "R", "PG-13"
// use in conjunction with region
func (b *DiscoverMoviesBuilder) CertificationLTE(c string) *DiscoverMoviesBuilder {
	b.opts.CertificationLTE = &c
	return b
}

// CertificationCountry sets the certification_country parameter. e.g. "US", "FR"
// use in conjunction with the certification, certification.gte and certification.lte filters
func (b *DiscoverMoviesBuilder) CertificationCountry(c string) *DiscoverMoviesBuilder {
	b.opts.CertificationCountry = &c
	return b
}

// IncludeVideo sets the include_video parameter. e.g. true, false (default: false)
func (b *DiscoverMoviesBuilder) IncludeVideo(v bool) *DiscoverMoviesBuilder {
	b.opts.IncludeVideo = &v
	return b
}

func (b *DiscoverMoviesBuilder) WithoutCompanies(ids ...string) *DiscoverMoviesBuilder {
	b.opts.WithoutCompanies = append(b.opts.WithoutCompanies, ids...)
	return b
}

func (b *DiscoverMoviesBuilder) WithCast(ids ...string) *DiscoverMoviesBuilder {
	b.opts.WithCast = append(b.opts.WithCast, ids...)
	return b
}

func (b *DiscoverMoviesBuilder) WithCrew(ids ...string) *DiscoverMoviesBuilder {
	b.opts.WithCrew = append(b.opts.WithCrew, ids...)
	return b
}

func (b *DiscoverMoviesBuilder) WithPeople(ids ...string) *DiscoverMoviesBuilder {
	b.opts.WithPeople = append(b.opts.WithPeople, ids...)
	return b
}

func (b *DiscoverMoviesBuilder) WithReleaseType(rt int) *DiscoverMoviesBuilder {
	b.opts.WithReleaseType = &rt
	return b
}

func (b *DiscoverMoviesBuilder) VoteAverageGTE(f float64) *DiscoverMoviesBuilder {
	b.opts.VoteAverageGTE = &f
	return b
}

func (b *DiscoverMoviesBuilder) VoteAverageLTE(f float64) *DiscoverMoviesBuilder {
	b.opts.VoteAverageLTE = &f
	return b
}

func (b *DiscoverMoviesBuilder) VoteCountGTE(i int) *DiscoverMoviesBuilder {
	b.opts.VoteCountGTE = &i
	return b
}

func (b *DiscoverMoviesBuilder) VoteCountLTE(i int) *DiscoverMoviesBuilder {
	b.opts.VoteCountLTE = &i
	return b
}

func (b *DiscoverMoviesBuilder) WithRuntimeGTE(i int) *DiscoverMoviesBuilder {
	b.opts.WithRuntimeGTE = &i
	return b
}

func (b *DiscoverMoviesBuilder) WithRuntimeLTE(i int) *DiscoverMoviesBuilder {
	b.opts.WithRuntimeLTE = &i
	return b
}

// Exec performs the request and returns the response.
func (b *DiscoverMoviesBuilder) Exec() (*types.MoviePaginatedResults, error) {
	path := "/discover/movie"
	params := make(url.Values)
	res := new(types.MoviePaginatedResults)

	baseOpts, err := utils.StructToURLValues(b.BaseDiscoverBuilder.BaseOpts)
	if err != nil {
		return nil, fmt.Errorf("convert base opts: %w", err)
	}
	opts, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("convert movie opts: %w", err)
	}

	// Combining baseOpts and opts
	for k, values := range baseOpts {
		for _, v := range values {
			params.Add(k, v)
		}
	}

	for k, values := range opts {
		for _, v := range values {
			params.Add(k, v)
		}
	}

	fmt.Println("params: ", params)

	err = b.client.DoRequest("GET", path, params, nil, res)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return res, nil
}

// DiscoverTVBuilder builds /discover/tv requests.
type DiscoverTVBuilder struct {
	client *client.Client
	*BaseDiscoverBuilder[*DiscoverTVBuilder]
	opts struct {
		FirstAirDateYear         *int     `url:"first_air_date_year,omitempty"`
		FirstAirDateGTE          *string  `url:"first_air_date.gte,omitempty"`
		FirstAirDateLTE          *string  `url:"first_air_date.lte,omitempty"`
		AirDateGTE               *string  `url:"air_date.gte,omitempty"`
		AirDateLTE               *string  `url:"air_date.lte,omitempty"`
		IncludeNullFirstAirDates *bool    `url:"include_null_first_air_dates,omitempty"`
		WithNetworks             []string `url:"with_networks,omitempty,pipe"`
		WithStatus               *string  `url:"with_status,omitempty"`
		WithType                 *string  `url:"with_type,omitempty"`
		VoteAverageGTE           *float64 `url:"vote_average.gte,omitempty"`
		VoteCountGTE             *int     `url:"vote_count.gte,omitempty"`
		WithRuntimeGTE           *int     `url:"with_runtime.gte,omitempty"`
		WithRuntimeLTE           *int     `url:"with_runtime.lte,omitempty"`
	}
}

func (b *DiscoverTVBuilder) isBuilder() {}

// NewDiscoverTVBuilder initializes the TV discover builder.
func NewDiscoverTVBuilder(c *client.Client) *DiscoverTVBuilder {
	b := &DiscoverTVBuilder{}
	b.client = c
	b.BaseDiscoverBuilder = &BaseDiscoverBuilder[*DiscoverTVBuilder]{self: b}
	return b
}

// TV-specific setters
func (b *DiscoverTVBuilder) FirstAirDateYear(y int) *DiscoverTVBuilder {
	b.opts.FirstAirDateYear = &y
	return b
}
func (b *DiscoverTVBuilder) FirstAirDateGTE(t time.Time) *DiscoverTVBuilder {
	s := t.Format("2006-01-02")
	b.opts.FirstAirDateGTE = &s
	return b
}
func (b *DiscoverTVBuilder) FirstAirDateLTE(t time.Time) *DiscoverTVBuilder {
	s := t.Format("2006-01-02")
	b.opts.FirstAirDateLTE = &s
	return b
}
func (b *DiscoverTVBuilder) AirDateGTE(t time.Time) *DiscoverTVBuilder {
	s := t.Format("2006-01-02")
	b.opts.AirDateGTE = &s
	return b
}
func (b *DiscoverTVBuilder) AirDateLTE(t time.Time) *DiscoverTVBuilder {
	s := t.Format("2006-01-02")
	b.opts.AirDateLTE = &s
	return b
}
func (b *DiscoverTVBuilder) IncludeNullFirstAirDates(v bool) *DiscoverTVBuilder {
	b.opts.IncludeNullFirstAirDates = &v
	return b
}

func (b *DiscoverTVBuilder) WithNetworks(ids ...string) *DiscoverTVBuilder {
	b.opts.WithNetworks = append(b.opts.WithNetworks, ids...)
	return b
}

func (b *DiscoverTVBuilder) WithStatus(s string) *DiscoverTVBuilder { b.opts.WithStatus = &s; return b }

func (b *DiscoverTVBuilder) WithType(t string) *DiscoverTVBuilder { b.opts.WithType = &t; return b }

func (b *DiscoverTVBuilder) VoteAverageGTE(f float64) *DiscoverTVBuilder {
	b.opts.VoteAverageGTE = &f
	return b
}

func (b *DiscoverTVBuilder) VoteCountGTE(i int) *DiscoverTVBuilder {
	b.opts.VoteCountGTE = &i
	return b
}

func (b *DiscoverTVBuilder) WithRuntimeGTE(i int) *DiscoverTVBuilder {
	b.opts.WithRuntimeGTE = &i
	return b
}

func (b *DiscoverTVBuilder) WithRuntimeLTE(i int) *DiscoverTVBuilder {
	b.opts.WithRuntimeLTE = &i
	return b
}

// Exec performs the request and returns the response.
func (b *DiscoverTVBuilder) Exec() (*types.TVShowPaginatedResults, error) {
	path := "/discover/tv"
	params := make(url.Values)
	res := new(types.TVShowPaginatedResults)

	baseOpts, err := utils.StructToURLValues(b.BaseDiscoverBuilder.BaseOpts)
	if err != nil {
		return nil, fmt.Errorf("convert base opts: %w", err)
	}
	opts, err := utils.StructToURLValues(b.opts)
	if err != nil {
		return nil, fmt.Errorf("convert movie opts: %w", err)
	}

	// Combining baseOpts and opts
	for k, values := range baseOpts {
		for _, v := range values {
			params.Add(k, v)
		}
	}

	for k, values := range opts {
		for _, v := range values {
			params.Add(k, v)
		}
	}

	fmt.Println("params: ", params)

	err = b.client.DoRequest("GET", path, params, nil, res)
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	return res, nil
}
