package types

// TVSeasonDetails represents the detailed information for a specific TV season.
// Includes fields for append_to_response results when requested.
// See: https://developer.themoviedb.org/reference/tv-season-details
type TVSeasonDetails struct {
	IDString     string                `json:"_id"` // Internal ID string
	AirDate      *string               `json:"air_date"` // Nullable, Format YYYY-MM-DD
	Episodes     []TVEpisodeListResult `json:"episodes"`
	Name         string                `json:"name"`
	Overview     string                `json:"overview"`
	ID           int                   `json:"id"` // TMDB season ID
	PosterPath   *string               `json:"poster_path"` // Nullable
	SeasonNumber int                   `json:"season_number"`
	VoteAverage  float64               `json:"vote_average"`

	// --- Appended fields --- 
	AccountStates    *TVSeasonAccountStatesResponse `json:"account_states,omitempty"` // Different structure than movie/tv
	AggregateCredits *AggregateCreditsResponse    `json:"aggregate_credits,omitempty"`
	Credits          *Credits                     `json:"credits,omitempty"`
	ExternalIDs      *TVSeasonExternalIDs         `json:"external_ids,omitempty"`
	Images           *TVSeasonImagesResponse      `json:"images,omitempty"`
	Translations     *TVSeasonTranslationsResponse `json:"translations,omitempty"`
	Videos           *TVSeasonVideosResponse      `json:"videos,omitempty"`
	WatchProviders   *WatchProviderResponse       `json:"watch/providers,omitempty"`
}

// TVSeasonListResult represents a TV season in a list (e.g., within Find results).
type TVSeasonListResult struct {
	AirDate      *string `json:"air_date"`
	EpisodeCount int     `json:"episode_count"`
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	Overview     string  `json:"overview"`
	PosterPath   *string `json:"poster_path"`
	SeasonNumber int     `json:"season_number"`
	ShowID       int     `json:"show_id"` // Added for context from Find endpoint
	VoteAverage  float64 `json:"vote_average"`
}

// EpisodeAccountState represents the rating status for a single episode within a season.
// Used in TVSeasonAccountStatesResponse.
type EpisodeAccountState struct {
	ID            int         `json:"id"`             // Episode ID
	EpisodeNumber int         `json:"episode_number"`
	Rated         interface{} `json:"rated"` // Can be boolean (false) or RatedInfo object
}

// TVSeasonAccountStatesResponse holds the account states for episodes within a season.
// Note the structure: an array of episode states.
// See: https://developer.themoviedb.org/reference/tv-season-account-states
type TVSeasonAccountStatesResponse struct {
	ID      int                   `json:"id"` // Season ID
	Results []EpisodeAccountState `json:"results"`
}

// TVSeasonChangesResponse represents the changes for a specific TV season.
// See: https://developer.themoviedb.org/reference/tv-season-changes
type TVSeasonChangesResponse struct {
	Changes []ChangeGroup `json:"changes"` // Uses common ChangeGroup
}

// TVSeasonExternalIDs represents external IDs specifically for a TV season.
// See: https://developer.themoviedb.org/reference/tv-season-external-ids
type TVSeasonExternalIDs struct {
	ID          int     `json:"id"` // Season ID
	FreebaseMID *string `json:"freebase_mid,omitempty"`
	FreebaseID  *string `json:"freebase_id,omitempty"`
	TVDBID      *int    `json:"tvdb_id,omitempty"`
	TVRageID    *int    `json:"tvrage_id,omitempty"` // Deprecated?
	WikidataID  *string `json:"wikidata_id,omitempty"`
}

// TVSeasonImagesResponse holds poster images for a TV season.
// See: https://developer.themoviedb.org/reference/tv-season-images
type TVSeasonImagesResponse struct {
	ID      int     `json:"id"` // Season ID
	Posters []Image `json:"posters"`
}

// TVSeasonTranslationData holds the translatable fields for a TV season.
type TVSeasonTranslationData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

// TVSeasonTranslation represents a translation for a TV season.
type TVSeasonTranslation struct {
	ISO3166_1   string                  `json:"iso_3166_1"`
	ISO639_1    string                  `json:"iso_639_1"`
	Name        string                  `json:"name"`
	EnglishName string                  `json:"english_name"`
	Data        TVSeasonTranslationData `json:"data"`
}

// TVSeasonTranslationsResponse holds translations for a TV season.
// See: https://developer.themoviedb.org/reference/tv-season-translations
type TVSeasonTranslationsResponse struct {
	ID           int                   `json:"id"` // Season ID
	Translations []TVSeasonTranslation `json:"translations"`
}

// TVSeasonVideosResponse holds videos for a TV season.
// See: https://developer.themoviedb.org/reference/tv-season-videos
type TVSeasonVideosResponse struct {
	ID      int     `json:"id"` // Season ID
	Results []Video `json:"results"`
}

