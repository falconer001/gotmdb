package types

// BelongsToCollection represents basic collection info embedded in MovieDetails.
type BelongsToCollection struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	PosterPath   *string `json:"poster_path"`   // Nullable
	BackdropPath *string `json:"backdrop_path"` // Nullable
}

// MovieDetails represents the detailed information for a movie.
// Includes fields for append_to_response results when requested.
// See: https://developer.themoviedb.org/reference/movie-details
type MovieDetails struct {
	Adult               bool                 `json:"adult"`
	BackdropPath        *string              `json:"backdrop_path"`                   // Nullable
	BelongsToCollection *BelongsToCollection `json:"belongs_to_collection,omitempty"` // Nullable
	Budget              int                  `json:"budget"`
	Genres              []Genre              `json:"genres"`   // Uses common Genre
	Homepage            *string              `json:"homepage"` // Nullable
	ID                  int                  `json:"id"`
	IMDbID              *string              `json:"imdb_id"` // Nullable, format tt1234567
	OriginCountry       []string             `json:"origin_country"`
	OriginalLanguage    string               `json:"original_language"`
	OriginalTitle       string               `json:"original_title"`
	Overview            string               `json:"overview"`
	Popularity          float64              `json:"popularity"`
	PosterPath          *string              `json:"poster_path"`          // Nullable
	ProductionCompanies []ProductionCompany  `json:"production_companies"` // Uses common ProductionCompany
	ProductionCountries []ProductionCountry  `json:"production_countries"` // Uses common ProductionCountry
	ReleaseDate         string               `json:"release_date"`         // Format YYYY-MM-DD
	Revenue             int64                `json:"revenue"`              // Use int64 for potentially large numbers
	Runtime             *int                 `json:"runtime"`              // Nullable, in minutes
	SpokenLanguages     []SpokenLanguage     `json:"spoken_languages"`     // Uses common SpokenLanguage
	Status              string               `json:"status"`               // e.g., "Rumored", "Planned", "In Production", "Post Production", "Released", "Canceled"
	Tagline             *string              `json:"tagline"`              // Nullable
	Title               string               `json:"title"`
	Video               bool                 `json:"video"` // Deprecated, use Videos append
	VoteAverage         float64              `json:"vote_average"`
	VoteCount           int                  `json:"vote_count"`

	// --- Appended fields ---
	AccountStates     *AccountState              `json:"account_states,omitempty"`
	AlternativeTitles *AlternativeTitlesResponse `json:"alternative_titles,omitempty"`
	Changes           *ItemChangesResponse       `json:"changes,omitempty"`      // Use ItemChangesResponse from changes.go
	Credits           *Credits                   `json:"credits,omitempty"`      // Uses common Credits
	ExternalIDs       *ExternalIDs               `json:"external_ids,omitempty"` // Uses common ExternalIDs
	Images            *ImageList                 `json:"images,omitempty"`       // Uses common ImageList
	Keywords          *KeywordsResponse          `json:"keywords,omitempty"`     // Uses KeywordsResponse from keywords.go
	Lists             *ListPaginatedResults      `json:"lists,omitempty"`        // Uses ListPaginatedResults from lists.go
	Recommendations   *MoviePaginatedResults     `json:"recommendations,omitempty"`
	ReleaseDates      *ReleaseDatesResponse      `json:"release_dates,omitempty"`
	Reviews           *ReviewPaginatedResults    `json:"reviews,omitempty"` // Uses ReviewPaginatedResults from common.go
	Similar           *MoviePaginatedResults     `json:"similar,omitempty"`
	Translations      *TranslationsResponse      `json:"translations,omitempty"`    // Uses common TranslationsResponse
	Videos            *VideoList                 `json:"videos,omitempty"`          // Uses common VideoList
	WatchProviders    *WatchProviderResponse     `json:"watch/providers,omitempty"` // Uses WatchProviderResponse from watch_providers.go
}

// MovieListResult represents a movie in a list (search results, popular, etc.).
type MovieListResult struct {
	Adult            bool    `json:"adult"`
	BackdropPath     *string `json:"backdrop_path"` // Nullable
	GenreIDs         []int   `json:"genre_ids"`
	ID               int     `json:"id"`
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    string  `json:"original_title"`
	Overview         string  `json:"overview"`
	Popularity       float64 `json:"popularity"`
	PosterPath       *string `json:"poster_path"`  // Nullable
	ReleaseDate      string  `json:"release_date"` // Format YYYY-MM-DD, can be empty
	Title            string  `json:"title"`
	Video            bool    `json:"video"` // Deprecated
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
}

// MoviePaginatedResults represents paginated movie list results.
type MoviePaginatedResults struct {
	Paginated                   // Embed common pagination fields
	Results   []MovieListResult `json:"results"`
}

// ReleaseDateInfo represents a single release date entry for a country.
type ReleaseDateInfo struct {
	Certification string   `json:"certification"`
	Descriptors   []string `json:"descriptors"`  // e.g., ["Premiere"]
	ISO639_1      *string  `json:"iso_639_1"`    // Nullable language code
	Note          *string  `json:"note"`         // Nullable note
	ReleaseDate   string   `json:"release_date"` // Timestamp (YYYY-MM-DDTHH:MM:SS.mmmZ)
	Type          int      `json:"type"`         // 1: Premiere, 2: Theatrical (limited), 3: Theatrical, 4: Digital, 5: Physical, 6: TV
}

// CountryReleaseDates holds all release dates for a specific country.
type CountryReleaseDates struct {
	ISO3166_1    string            `json:"iso_3166_1"` // Country code
	ReleaseDates []ReleaseDateInfo `json:"release_dates"`
}

// ReleaseDatesResponse holds the release dates grouped by country.
// See: https://developer.themoviedb.org/reference/movie-release-dates
type ReleaseDatesResponse struct {
	ID      int                   `json:"id"` // Movie ID
	Results []CountryReleaseDates `json:"results"`
}

// DateRange represents the min/max date range for Now Playing or Upcoming movies.
type DateRange struct {
	Maximum string `json:"maximum"` // Format YYYY-MM-DD
	Minimum string `json:"minimum"` // Format YYYY-MM-DD
}

// NowPlayingResponse represents the response for the now playing movies endpoint.
// See: https://developer.themoviedb.org/reference/movie-now-playing-list
type NowPlayingResponse struct {
	Dates     DateRange         `json:"dates"`
	Paginated                   // Embed common pagination fields
	Results   []MovieListResult `json:"results"`
}

// UpcomingResponse represents the response for the upcoming movies endpoint.
// See: https://developer.themoviedb.org/reference/movie-upcoming-list
type UpcomingResponse struct {
	Dates     DateRange         `json:"dates"`
	Paginated                   // Embed common pagination fields
	Results   []MovieListResult `json:"results"`
}

// RatingRequest is the request body for adding/updating a movie or TV rating.
type RatingRequest struct {
	Value float64 `json:"value"` // Rating value (0.5 to 10.0 in 0.5 increments)
}
