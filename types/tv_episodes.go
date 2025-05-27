package types

// TVEpisodeDetails represents the detailed information for a specific TV episode.
// Includes fields for append_to_response results when requested.
// See: https://developer.themoviedb.org/reference/tv-episode-details
type TVEpisodeDetails struct {
	AirDate       string       `json:"air_date"` // Format YYYY-MM-DD
	Crew          []CrewMember `json:"crew"` // Uses common CrewMember
	EpisodeNumber int          `json:"episode_number"`
	GuestStars    []CastMember `json:"guest_stars"` // Uses common CastMember
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	Overview      string       `json:"overview"`
	ProductionCode string      `json:"production_code"`
	Runtime       *int         `json:"runtime"` // Nullable
	SeasonNumber  int          `json:"season_number"`
	ShowID        int          `json:"show_id"`
	StillPath     *string      `json:"still_path"` // Nullable
	VoteAverage   float64      `json:"vote_average"`
	VoteCount     int          `json:"vote_count"`

	// --- Appended fields --- 
	AccountStates *EpisodeAccountStateSingle `json:"account_states,omitempty"`
	Credits       *TVEpisodeCreditsResponse  `json:"credits,omitempty"`
	ExternalIDs   *TVEpisodeExternalIDs    `json:"external_ids,omitempty"`
	Images        *TVEpisodeImagesResponse   `json:"images,omitempty"`
	Translations  *TVEpisodeTranslationsResponse `json:"translations,omitempty"`
	Videos        *TVEpisodeVideosResponse   `json:"videos,omitempty"`
}

// TVEpisodeListResult represents a TV episode in a list (e.g., within TVSeasonDetails or Find results).
type TVEpisodeListResult struct {
	AirDate       string       `json:"air_date"`
	EpisodeNumber int          `json:"episode_number"`
	ID            int          `json:"id"`
	Name          string       `json:"name"`
	Overview      string       `json:"overview"`
	ProductionCode string      `json:"production_code"`
	Runtime       *int         `json:"runtime"` // Nullable
	SeasonNumber  int          `json:"season_number"`
	ShowID        int          `json:"show_id"`
	StillPath     *string      `json:"still_path"` // Nullable
	VoteAverage   float64      `json:"vote_average"`
	VoteCount     int          `json:"vote_count"`
	Crew          []CrewMember `json:"crew,omitempty"` // Included in season details
	GuestStars    []CastMember `json:"guest_stars,omitempty"` // Included in season details
	Order         *int         `json:"order,omitempty"` // Included in episode group details
}

// EpisodeAccountStateSingle represents the account state for a single episode.
// See: https://developer.themoviedb.org/reference/tv-episode-account-states
type EpisodeAccountStateSingle struct {
	ID            int         `json:"id"` // Episode ID
	Rated         interface{} `json:"rated"` // boolean (false) or RatedInfo object
	EpisodeNumber int         `json:"episode_number"` // Included for context, though ID is primary key
}

// TVEpisodeChangesResponse represents the changes for a specific TV episode.
// See: https://developer.themoviedb.org/reference/tv-episode-changes
type TVEpisodeChangesResponse struct {
	Changes []ChangeGroup `json:"changes"` // Uses common ChangeGroup
}

// TVEpisodeCreditsResponse holds the credits (cast, crew, guest stars) for a TV episode.
// See: https://developer.themoviedb.org/reference/tv-episode-credits
type TVEpisodeCreditsResponse struct {
	ID         int          `json:"id"` // Episode ID
	Cast       []CastMember `json:"cast"`
	Crew       []CrewMember `json:"crew"`
	GuestStars []CastMember `json:"guest_stars"`
}

// TVEpisodeExternalIDs represents external IDs specifically for a TV episode.
// See: https://developer.themoviedb.org/reference/tv-episode-external-ids
type TVEpisodeExternalIDs struct {
	ID          int     `json:"id"` // Episode ID
	IMDbID      *string `json:"imdb_id,omitempty"`
	FreebaseMID *string `json:"freebase_mid,omitempty"`
	FreebaseID  *string `json:"freebase_id,omitempty"`
	TVDBID      *int    `json:"tvdb_id,omitempty"`
	TVRageID    *int    `json:"tvrage_id,omitempty"` // Deprecated?
	WikidataID  *string `json:"wikidata_id,omitempty"`
}

// TVEpisodeImagesResponse holds still images for a TV episode.
// See: https://developer.themoviedb.org/reference/tv-episode-images
type TVEpisodeImagesResponse struct {
	ID     int     `json:"id"` // Episode ID
	Stills []Image `json:"stills"`
}

// TVEpisodeTranslationData holds the translatable fields for a TV episode.
type TVEpisodeTranslationData struct {
	Name     string `json:"name"`
	Overview string `json:"overview"`
}

// TVEpisodeTranslation represents a translation for a TV episode.
type TVEpisodeTranslation struct {
	ISO3166_1   string                   `json:"iso_3166_1"`
	ISO639_1    string                   `json:"iso_639_1"`
	Name        string                   `json:"name"`
	EnglishName string                   `json:"english_name"`
	Data        TVEpisodeTranslationData `json:"data"`
}

// TVEpisodeTranslationsResponse holds translations for a TV episode.
// See: https://developer.themoviedb.org/reference/tv-episode-translations
type TVEpisodeTranslationsResponse struct {
	ID           int                    `json:"id"` // Episode ID
	Translations []TVEpisodeTranslation `json:"translations"`
}

// TVEpisodeVideosResponse holds videos for a TV episode.
// See: https://developer.themoviedb.org/reference/tv-episode-videos
type TVEpisodeVideosResponse struct {
	ID      int     `json:"id"` // Episode ID
	Results []Video `json:"results"`
}

