package types

// MediaInfo represents basic media information within a CreditDetails response.
// It combines fields common to movies and TV shows.
type MediaInfo struct {
	ID                 int      `json:"id"`
	Name               string   `json:"name,omitempty"` // TV Show Name
	OriginalName       string   `json:"original_name,omitempty"` // TV Show Original Name
	Title              string   `json:"title,omitempty"` // Movie Title
	OriginalTitle      string   `json:"original_title,omitempty"` // Movie Original Title
	Character          string   `json:"character,omitempty"` // Character name (if credit_type is cast)
	PosterPath         *string  `json:"poster_path"` // Nullable
	VoteAverage        float64  `json:"vote_average"`
	VoteCount          int      `json:"vote_count"`
	Overview           string   `json:"overview"`
	ReleaseDate        string   `json:"release_date,omitempty"` // Movie release date (YYYY-MM-DD)
	FirstAirDate       string   `json:"first_air_date,omitempty"` // TV first air date (YYYY-MM-DD)
	Adult              *bool    `json:"adult,omitempty"` // Movie adult status (nullable)
	BackdropPath       *string  `json:"backdrop_path"` // Nullable
	GenreIDs           []int    `json:"genre_ids"`
	OriginalLanguage   string   `json:"original_language"`
	Popularity         float64  `json:"popularity"`
	Video              *bool    `json:"video,omitempty"` // Movie video status (nullable)
	OriginCountry      []string `json:"origin_country,omitempty"` // TV origin countries
	EpisodeCount       *int     `json:"episode_count,omitempty"` // TV episode count (if applicable)
}

// PersonInfo represents basic person information within a CreditDetails response.
type PersonInfo struct {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	OriginalName string  `json:"original_name"`
	ProfilePath  *string `json:"profile_path"` // Nullable
	Gender       int     `json:"gender"` // 0: Not set, 1: Female, 2: Male, 3: Non-binary
	Adult        bool    `json:"adult"`
}

// CreditDetails represents the details of a specific credit ID.
// See: https://developer.themoviedb.org/reference/credit-details
type CreditDetails struct {
	CreditType string     `json:"credit_type"` // "cast" or "crew"
	Department string     `json:"department"`   // Department (if crew)
	Job        string     `json:"job"`         // Job (if crew)
	Media      MediaInfo  `json:"media"`       // Details of the movie or TV show
	MediaType  string     `json:"media_type"`  // "movie" or "tv"
	ID         string     `json:"id"`         // The credit ID itself
	Person     PersonInfo `json:"person"`      // Details of the person
}

