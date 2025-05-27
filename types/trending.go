package types

// TrendingResult represents a single item in a trending results list.
// The structure is similar to SearchMultiResult, varying based on media_type.
// See: https://developer.themoviedb.org/reference/trending-all
type TrendingResult struct {
	Adult            *bool    `json:"adult,omitempty"` // Movie/Person
	BackdropPath     *string  `json:"backdrop_path,omitempty"` // Movie/TV
	ID               int      `json:"id"`
	Title            *string  `json:"title,omitempty"` // Movie
	OriginalLanguage *string  `json:"original_language,omitempty"` // Movie/TV
	OriginalTitle    *string  `json:"original_title,omitempty"` // Movie
	Overview         *string  `json:"overview,omitempty"` // Movie/TV
	PosterPath       *string  `json:"poster_path,omitempty"` // Movie/TV/Person
	MediaType        string   `json:"media_type"` // "movie", "tv", "person"
	GenreIDs         []int    `json:"genre_ids,omitempty"` // Movie/TV
	Popularity       float64  `json:"popularity"`
	ReleaseDate      *string  `json:"release_date,omitempty"` // Movie (YYYY-MM-DD)
	Video            *bool    `json:"video,omitempty"` // Movie
	VoteAverage      *float64 `json:"vote_average,omitempty"` // Movie/TV
	VoteCount        *int     `json:"vote_count,omitempty"` // Movie/TV
	Name             *string  `json:"name,omitempty"` // TV/Person
	OriginalName     *string  `json:"original_name,omitempty"` // TV/Person
	FirstAirDate     *string  `json:"first_air_date,omitempty"` // TV (YYYY-MM-DD)
	OriginCountry    []string `json:"origin_country,omitempty"` // TV
	Gender           *int     `json:"gender,omitempty"` // Person
	KnownForDepartment *string `json:"known_for_department,omitempty"` // Person
	ProfilePath      *string  `json:"profile_path,omitempty"` // Person
	KnownFor         []KnownForItem `json:"known_for,omitempty"` // Person (Uses KnownForItem from people.go)
}

// TrendingResponse represents paginated trending results.
type TrendingResponse struct {
	Paginated         // Embed common pagination fields
	Results   []TrendingResult `json:"results"`
}

