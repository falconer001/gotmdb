package types

// SearchMultiResult represents a single result in a multi-search response.
// The structure varies depending on the media_type (movie, tv, person).
// Using json.RawMessage allows flexible handling, or define a struct with all possible fields.
// For a type-safe approach, embed common fields and include specific ones with omitempty.
type SearchMultiResult struct {
	Adult              *bool          `json:"adult,omitempty"`         // Movie/Person
	BackdropPath       *string        `json:"backdrop_path,omitempty"` // Movie/TV
	ID                 int            `json:"id"`
	Title              *string        `json:"title,omitempty"`             // Movie
	OriginalLanguage   *string        `json:"original_language,omitempty"` // Movie/TV
	OriginalTitle      *string        `json:"original_title,omitempty"`    // Movie
	Overview           *string        `json:"overview,omitempty"`          // Movie/TV
	PosterPath         *string        `json:"poster_path,omitempty"`       // Movie/TV/Person
	MediaType          string         `json:"media_type"`                  // "movie", "tv", "person"
	GenreIDs           []int          `json:"genre_ids,omitempty"`         // Movie/TV
	Popularity         float64        `json:"popularity"`
	ReleaseDate        *string        `json:"release_date,omitempty"`         // Movie (YYYY-MM-DD)
	Video              *bool          `json:"video,omitempty"`                // Movie
	VoteAverage        *float64       `json:"vote_average,omitempty"`         // Movie/TV
	VoteCount          *int           `json:"vote_count,omitempty"`           // Movie/TV
	Name               *string        `json:"name,omitempty"`                 // TV/Person
	OriginalName       *string        `json:"original_name,omitempty"`        // TV/Person
	FirstAirDate       *string        `json:"first_air_date,omitempty"`       // TV (YYYY-MM-DD)
	OriginCountry      []string       `json:"origin_country,omitempty"`       // TV
	Gender             *int           `json:"gender,omitempty"`               // Person
	KnownForDepartment *string        `json:"known_for_department,omitempty"` // Person
	ProfilePath        *string        `json:"profile_path,omitempty"`         // Person
	KnownFor           []KnownForItem `json:"known_for,omitempty"`            // Person (Uses KnownForItem from people.go)
}

// SearchMultiResponse represents paginated multi-search results.
// See: https://developer.themoviedb.org/reference/search-multi
type SearchMultiResponse struct {
	Paginated                     // Embed common pagination fields
	Results   []SearchMultiResult `json:"results"`
}

// Note: Other search result types (CompanySearchResult, CollectionSearchResult,
// KeywordSearchResponse, MoviePaginatedResults, PersonPaginatedResults,
// TVShowPaginatedResults) are defined in their respective files (companies.go,
// collections.go, keywords.go, movies.go, people.go, tv.go).
