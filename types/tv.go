package types

// LastEpisodeToAir represents basic info about the last episode aired.
// Used within TVDetails.
type LastEpisodeToAir struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Overview      string  `json:"overview"`
	VoteAverage   float64 `json:"vote_average"`
	VoteCount     int     `json:"vote_count"`
	AirDate       string  `json:"air_date"` // Format YYYY-MM-DD
	EpisodeNumber int     `json:"episode_number"`
	ProductionCode string `json:"production_code"`
	Runtime       *int    `json:"runtime"` // Nullable
	SeasonNumber  int     `json:"season_number"`
	ShowID        int     `json:"show_id"`
	StillPath     *string `json:"still_path"` // Nullable
}

// NextEpisodeToAir represents basic info about the next episode to air.
// Used within TVDetails.
type NextEpisodeToAir struct {
	ID            int     `json:"id"`
	Name          string  `json:"name"`
	Overview      string  `json:"overview"`
	VoteAverage   float64 `json:"vote_average"` // Often 0 for unaired
	VoteCount     int     `json:"vote_count"`     // Often 0 for unaired
	AirDate       string  `json:"air_date"` // Format YYYY-MM-DD
	EpisodeNumber int     `json:"episode_number"`
	ProductionCode string `json:"production_code"`
	Runtime       *int    `json:"runtime"` // Nullable
	SeasonNumber  int     `json:"season_number"`
	ShowID        int     `json:"show_id"`
	StillPath     *string `json:"still_path"` // Nullable
}

// TVSeason represents basic info about a TV season.
// Used within TVDetails.
type TVSeason struct {
	AirDate      *string `json:"air_date"` // Nullable, Format YYYY-MM-DD
	EpisodeCount int     `json:"episode_count"`
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"` // Nullable
	SeasonNumber int     `json:"season_number"`
	VoteAverage  float64 `json:"vote_average"`
}

// TVDetails represents the detailed information for a TV show.
// Includes fields for append_to_response results when requested.
// See: https://developer.themoviedb.org/reference/tv-series-details
type TVDetails struct {
	Adult            bool              `json:"adult"` // Typically false for TV
	BackdropPath     *string           `json:"backdrop_path"` // Nullable
	CreatedBy        []Creator         `json:"created_by"`
	EpisodeRunTime   []int             `json:"episode_run_time"`
	FirstAirDate     string            `json:"first_air_date"` // Format YYYY-MM-DD
	Genres           []Genre           `json:"genres"` // Uses common Genre
	Homepage         *string           `json:"homepage"` // Nullable
	ID               int               `json:"id"`
	InProduction     bool              `json:"in_production"`
	Languages        []string          `json:"languages"` // List of ISO 639-1 codes
	LastAirDate      string            `json:"last_air_date"` // Format YYYY-MM-DD
	LastEpisodeToAir *LastEpisodeToAir `json:"last_episode_to_air,omitempty"` // Nullable
	Name             string            `json:"name"`
	NextEpisodeToAir *NextEpisodeToAir `json:"next_episode_to_air,omitempty"` // Nullable
	Networks         []Network         `json:"networks"` // Uses Network from networks.go
	NumberOfEpisodes int               `json:"number_of_episodes"`
	NumberOfSeasons  int               `json:"number_of_seasons"`
	OriginCountry    []string          `json:"origin_country"` // List of ISO 3166-1 codes
	OriginalLanguage string            `json:"original_language"`
	OriginalName     string            `json:"original_name"`
	Overview         string            `json:"overview"`
	Popularity       float64           `json:"popularity"`
	PosterPath       *string           `json:"poster_path"` // Nullable
	ProductionCompanies []ProductionCompany `json:"production_companies"` // Uses common ProductionCompany
	ProductionCountries []ProductionCountry `json:"production_countries"` // Uses common ProductionCountry
	Seasons          []TVSeason        `json:"seasons"`
	SpokenLanguages  []SpokenLanguage  `json:"spoken_languages"` // Uses common SpokenLanguage
	Status           string            `json:"status"` // e.g., "Returning Series", "Ended", "Canceled"
	Tagline          *string           `json:"tagline"` // Nullable
	Type             string            `json:"type"` // e.g., "Scripted", "Reality"
	VoteAverage      float64           `json:"vote_average"`
	VoteCount        int               `json:"vote_count"`

	// --- Appended fields --- 
	AccountStates       *AccountState             `json:"account_states,omitempty"`
	AggregateCredits    *AggregateCreditsResponse `json:"aggregate_credits,omitempty"`
	AlternativeTitles   *AlternativeTitlesResponse `json:"alternative_titles,omitempty"`
	Changes             *ItemChangesResponse      `json:"changes,omitempty"`
	ContentRatings      *ContentRatingsResponse   `json:"content_ratings,omitempty"`
	Credits             *Credits                  `json:"credits,omitempty"`
	EpisodeGroups       *EpisodeGroupsResponse    `json:"episode_groups,omitempty"`
	ExternalIDs         *TVExternalIDs            `json:"external_ids,omitempty"`
	Images              *TVImagesResponse         `json:"images,omitempty"`
	Keywords            *KeywordsResponse         `json:"keywords,omitempty"` // API docs say results, check response
	Recommendations     *TVShowPaginatedResults   `json:"recommendations,omitempty"`
	Reviews             *ReviewPaginatedResults   `json:"reviews,omitempty"`
	ScreenedTheatrically *ScreenedTheatricallyResponse `json:"screened_theatrically,omitempty"`
	Similar             *TVShowPaginatedResults   `json:"similar,omitempty"`
	Translations        *TVTranslationsResponse   `json:"translations,omitempty"`
	Videos              *VideoList                `json:"videos,omitempty"`
	WatchProviders      *WatchProviderResponse    `json:"watch/providers,omitempty"`
}

// Creator represents a creator of a TV show.
type Creator struct {
	ID          int     `json:"id"`
	CreditID    string  `json:"credit_id"`
	Name        string  `json:"name"`
	Gender      int     `json:"gender"`
	ProfilePath *string `json:"profile_path"` // Nullable
}

// TVListResult represents a TV show in a list (search results, popular, etc.).
type TVListResult struct {
	Adult            bool     `json:"adult"` // Typically false
	BackdropPath     *string  `json:"backdrop_path"` // Nullable
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"`
	OriginCountry    []string `json:"origin_country"`
	OriginalLanguage string   `json:"original_language"`
	OriginalName     string   `json:"original_name"`
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"` // Nullable
	FirstAirDate     string   `json:"first_air_date"` // Format YYYY-MM-DD, can be empty
	Name             string   `json:"name"`
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
}

// TVShowPaginatedResults represents paginated TV show list results.
type TVShowPaginatedResults struct {
	Paginated         // Embed common pagination fields
	Results   []TVListResult `json:"results"`
}

// Role represents a specific role played by an actor in AggregateCredits.
type Role struct {
	CreditID    string `json:"credit_id"`
	Character   string `json:"character"`
	EpisodeCount int    `json:"episode_count"`
}

// AggregateCastMember represents a cast member in aggregate credits.
type AggregateCastMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	Roles              []Role  `json:"roles"`
	TotalEpisodeCount  int     `json:"total_episode_count"`
	Order              int     `json:"order"` // Order in the main cast list
}

// Job represents a specific job held by a crew member in AggregateCredits.
type Job struct {
	CreditID    string `json:"credit_id"`
	Job         string `json:"job"`
	EpisodeCount int    `json:"episode_count"`
}

// AggregateCrewMember represents a crew member in aggregate credits.
type AggregateCrewMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`
	ID                 int     `json:"id"`
	KnownForDepartment string  `json:"known_for_department"`
	Name               string  `json:"name"`
	OriginalName       string  `json:"original_name"`
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"`
	Jobs               []Job   `json:"jobs"`
	Department         string  `json:"department"`
	TotalEpisodeCount  int     `json:"total_episode_count"`
}

// AggregateCreditsResponse holds the aggregate cast and crew for a TV show.
// See: https://developer.themoviedb.org/reference/tv-series-aggregate-credits
type AggregateCreditsResponse struct {
	ID   int                   `json:"id"` // TV Show ID
	Cast []AggregateCastMember `json:"cast"`
	Crew []AggregateCrewMember `json:"crew"`
}

// ContentRating represents a content rating for a specific country.
type ContentRating struct {
	Descriptors []string `json:"descriptors"` // Optional descriptors
	ISO3166_1   string   `json:"iso_3166_1"`   // Country code
	Rating      string   `json:"rating"`      // e.g., "TV-MA"
}

// ContentRatingsResponse holds the content ratings for a TV show.
// See: https://developer.themoviedb.org/reference/tv-series-content-ratings
type ContentRatingsResponse struct {
	ID      int             `json:"id"` // TV Show ID
	Results []ContentRating `json:"results"`
}

// EpisodeGroup represents a group of episodes (e.g., a special collection).
type EpisodeGroup struct {
	Description  string   `json:"description"`
	EpisodeCount int      `json:"episode_count"`
	GroupCount   int      `json:"group_count"` // Number of subgroups/episodes within this group
	ID           string   `json:"id"`         // Group ID (string)
	Name         string   `json:"name"`
	Network      *Network `json:"network,omitempty"` // Nullable, uses Network from networks.go
	Type         int      `json:"type"`         // Type identifier (e.g., 1 for original air date)
}

// EpisodeGroupsResponse holds the episode groups for a TV show.
// See: https://developer.themoviedb.org/reference/tv-series-episode-groups
type EpisodeGroupsResponse struct {
	ID           int            `json:"id"` // TV Show ID
	Results      []EpisodeGroup `json:"results"`
	TotalResults int            `json:"total_results"` // Note: Not paginated
}

// TVExternalIDs represents external IDs specifically for a TV show.
// See: https://developer.themoviedb.org/reference/tv-series-external-ids
type TVExternalIDs struct {
	ID          int     `json:"id"` // TV Show ID
	IMDbID      *string `json:"imdb_id,omitempty"`
	WikidataID  *string `json:"wikidata_id,omitempty"`
	FacebookID  *string `json:"facebook_id,omitempty"`
	InstagramID *string `json:"instagram_id,omitempty"`
	TwitterID   *string `json:"twitter_id,omitempty"`
	FreebaseMID *string `json:"freebase_mid,omitempty"` // Deprecated?
	FreebaseID  *string `json:"freebase_id,omitempty"`  // Deprecated?
	TVDBID      *int    `json:"tvdb_id,omitempty"`
	TVRageID    *int    `json:"tvrage_id,omitempty"`    // Deprecated?
}

// TVImagesResponse holds images for a TV show.
// See: https://developer.themoviedb.org/reference/tv-series-images
type TVImagesResponse struct {
	ID        int     `json:"id"` // TV Show ID
	Backdrops []Image `json:"backdrops"`
	Logos     []Image `json:"logos"`
	Posters   []Image `json:"posters"`
}

// ScreeningInfo represents an episode screened theatrically.
type ScreeningInfo struct {
	EpisodeNumber int `json:"episode_number"`
	SeasonNumber  int `json:"season_number"`
}

// ScreenedTheatricallyResponse holds episodes screened theatrically.
// See: https://developer.themoviedb.org/reference/tv-series-screened-theatrically
type ScreenedTheatricallyResponse struct {
	ID      int             `json:"id"` // TV Show ID
	Results []ScreeningInfo `json:"results"`
}

// TVTranslationData holds the translatable fields for a TV show.
type TVTranslationData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
	Homepage string `json:"homepage,omitempty"`
	Tagline  string `json:"tagline,omitempty"`
}

// TVTranslation represents a translation for a TV show.
type TVTranslation struct {
	ISO3166_1   string            `json:"iso_3166_1"`
	ISO639_1    string            `json:"iso_639_1"`
	Name        string            `json:"name"`
	EnglishName string            `json:"english_name"`
	Data        TVTranslationData `json:"data"`
}

// TVTranslationsResponse holds translations for a TV show.
// See: https://developer.themoviedb.org/reference/tv-series-translations
type TVTranslationsResponse struct {
	ID           int             `json:"id"` // TV Show ID
	Translations []TVTranslation `json:"translations"`
}

// TVVideosResponse holds videos for a TV show.
// See: https://developer.themoviedb.org/reference/tv-series-videos
type TVVideosResponse struct {
	ID      int     `json:"id"` // TV Show ID
	Results []Video `json:"results"`
}

