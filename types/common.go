package types

// Genre represents a movie or TV show genre.
// See: https://developer.themoviedb.org/reference/genre-movie-list
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// ProductionCompany represents a production company.
// Common field in MovieDetails, TVDetails, etc.
type ProductionCompany struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"` // Pointer for nullable
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

// ProductionCountry represents a production country.
// Common field in MovieDetails, TVDetails, etc.
type ProductionCountry struct {
	ISO3166_1 string `json:"iso_3166_1"` // e.g., "US"
	Name      string `json:"name"`       // e.g., "United States of America"
}

// SpokenLanguage represents a spoken language.
// Common field in MovieDetails, TVDetails, etc.
type SpokenLanguage struct {
	EnglishName string `json:"english_name"` // e.g., "English"
	ISO639_1    string `json:"iso_639_1"`    // e.g., "en"
	Name        string `json:"name"`         // e.g., "English"
}

// Video represents a video (trailer, teaser, etc.).
// Used in MovieVideos, TVVideos, etc., and append_to_response.
// See: https://developer.themoviedb.org/reference/movie-videos
type Video struct {
	ISO639_1    string `json:"iso_639_1"`    // Language code (e.g., "en")
	ISO3166_1   string `json:"iso_3166_1"`   // Country code (e.g., "US")
	Name        string `json:"name"`         // Video title
	Key         string `json:"key"`          // Platform key (e.g., YouTube ID)
	Site        string `json:"site"`         // Platform name (e.g., "YouTube")
	Size        int    `json:"size"`         // Video resolution (e.g., 1080)
	Type        string `json:"type"`         // Video type (e.g., "Trailer", "Teaser", "Featurette")
	Official    bool   `json:"official"`     // Whether the video is official
	PublishedAt string `json:"published_at"` // Timestamp (e.g., "2023-03-15T14:00:00.000Z")
	ID          string `json:"id"`           // Video ID (unique string)
}

// VideoList represents a list of videos, typically part of an append_to_response.
type VideoList struct {
	ID      int     `json:"id"`
	Results []Video `json:"results"`
}

// Image represents an image (poster, backdrop, logo, profile, still).
// Used in various image endpoints and append_to_response.
// See: https://developer.themoviedb.org/reference/movie-images
type Image struct {
	AspectRatio float64 `json:"aspect_ratio"` // Image aspect ratio
	Height      int     `json:"height"`       // Image height in pixels
	ISO639_1    *string `json:"iso_639_1"`    // Language code (nullable)
	FilePath    string  `json:"file_path"`    // Image path (append to base URL)
	VoteAverage float64 `json:"vote_average"` // Average vote score
	VoteCount   int     `json:"vote_count"`   // Number of votes
	Width       int     `json:"width"`        // Image width in pixels
}

// ImageList represents lists of posters, backdrops, and logos.
// Used in append_to_response for movies, TV shows, etc.
type ImageList struct {
	Backdrops []Image `json:"backdrops"`
	ID        int     `json:"id"`
	Logos     []Image `json:"logos"`
	Posters   []Image `json:"posters"`
}

// CastMember represents a cast member in credits.
// See: https://developer.themoviedb.org/reference/movie-credits
type CastMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`               // 0: Not set, 1: Female, 2: Male, 3: Non-binary
	ID                 int     `json:"id"`                   // Person ID
	KnownForDepartment string  `json:"known_for_department"` // e.g., "Acting"
	Name               string  `json:"name"`                 // Actor's name
	OriginalName       string  `json:"original_name"`        // Actor's original name
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"` // Actor's profile image path (nullable)
	CastID             int     `json:"cast_id"`      // ID specific to this cast credit
	Character          string  `json:"character"`    // Character name
	CreditID           string  `json:"credit_id"`    // Credit ID (unique string)
	Order              int     `json:"order"`        // Order in the credits list
}

// CrewMember represents a crew member in credits.
// See: https://developer.themoviedb.org/reference/movie-credits
type CrewMember struct {
	Adult              bool    `json:"adult"`
	Gender             int     `json:"gender"`               // 0: Not set, 1: Female, 2: Male, 3: Non-binary
	ID                 int     `json:"id"`                   // Person ID
	KnownForDepartment string  `json:"known_for_department"` // e.g., "Directing", "Writing"
	Name               string  `json:"name"`                 // Crew member's name
	OriginalName       string  `json:"original_name"`        // Crew member's original name
	Popularity         float64 `json:"popularity"`
	ProfilePath        *string `json:"profile_path"` // Crew member's profile image path (nullable)
	CreditID           string  `json:"credit_id"`    // Credit ID (unique string)
	Department         string  `json:"department"`   // e.g., "Directing", "Production"
	Job                string  `json:"job"`          // e.g., "Director", "Producer"
}

// Credits represents the cast and crew for a movie or TV show.
// Used in append_to_response and dedicated credit endpoints.
type Credits struct {
	ID   int          `json:"id"`
	Cast []CastMember `json:"cast"`
	Crew []CrewMember `json:"crew"`
}

// ExternalIDs represents common external IDs for movies, TV shows, or people.
// Fields may be nullable depending on the resource type.
// See: https://developer.themoviedb.org/reference/movie-external-ids
type ExternalIDs struct {
	ID          int     `json:"id"` // TMDB ID (present in some responses)
	IMDbID      *string `json:"imdb_id,omitempty"`
	WikidataID  *string `json:"wikidata_id,omitempty"`
	FacebookID  *string `json:"facebook_id,omitempty"`
	InstagramID *string `json:"instagram_id,omitempty"`
	TwitterID   *string `json:"twitter_id,omitempty"`
	TVDBID      *int    `json:"tvdb_id,omitempty"`      // TV specific
	TVRageID    *int    `json:"tvrage_id,omitempty"`    // TV/Person specific (deprecated?)
	FreebaseMID *string `json:"freebase_mid,omitempty"` // Deprecated?
	FreebaseID  *string `json:"freebase_id,omitempty"`  // Deprecated?
}

// Keyword represents a keyword associated with a movie or TV show.
// See: https://developer.themoviedb.org/reference/movie-keywords
type Keyword struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// KeywordsResponse is used for endpoints returning a list of keywords.
type KeywordsResponse struct {
	ID       int       `json:"id,omitempty"`       // Movie/TV ID, present on movie/tv keyword endpoints
	Keywords []Keyword `json:"keywords,omitempty"` // Field name varies (keywords/results), use omitempty
	Results  []Keyword `json:"results,omitempty"`  // Field name varies (keywords/results), use omitempty
}

// Review represents a user review.
// See: https://developer.themoviedb.org/reference/movie-reviews
type Review struct {
	Author        string        `json:"author"`
	AuthorDetails AuthorDetails `json:"author_details"`
	Content       string        `json:"content"`
	CreatedAt     string        `json:"created_at"` // Timestamp
	ID            string        `json:"id"`         // Review ID (unique string)
	UpdatedAt     string        `json:"updated_at"` // Timestamp
	URL           string        `json:"url"`
}

// AuthorDetails contains information about a review author.
type AuthorDetails struct {
	Name       string   `json:"name"`
	Username   string   `json:"username"`
	AvatarPath *string  `json:"avatar_path"` // Nullable
	Rating     *float64 `json:"rating"`      // Nullable (0.5-10.0)
}

// ReviewPaginatedResults represents paginated review results.
type ReviewPaginatedResults struct {
	ID           int      `json:"id"` // Movie/TV ID
	Page         int      `json:"page"`
	Results      []Review `json:"results"`
	TotalPages   int      `json:"total_pages"`
	TotalResults int      `json:"total_results"`
}

// TranslationData holds the translated text fields.
// Used within Translation structs.
type TranslationData struct {
	Title     string `json:"title,omitempty"`     // Movie
	Overview  string `json:"overview,omitempty"`  // Movie/TV/Collection/Season/Episode
	Homepage  string `json:"homepage,omitempty"`  // Movie/TV
	Tagline   string `json:"tagline,omitempty"`   // Movie/TV
	Name      string `json:"name,omitempty"`      // TV/Collection/Season/Episode
	Biography string `json:"biography,omitempty"` // Person
}

// Translation represents a single translation for a resource.
// See: https://developer.themoviedb.org/reference/movie-translations
type Translation struct {
	ISO3166_1   string          `json:"iso_3166_1"`   // Country code
	ISO639_1    string          `json:"iso_639_1"`    // Language code
	Name        string          `json:"name"`         // Language name in its own language
	EnglishName string          `json:"english_name"` // Language name in English
	Data        TranslationData `json:"data"`
}

// TranslationsResponse is a common structure for translation endpoints.
type TranslationsResponse struct {
	ID           int           `json:"id"` // ID of the translated resource (Movie, TV, Person, etc.)
	Translations []Translation `json:"translations"`
}

// AccountState represents the user's state for a movie or TV show.
// See: https://developer.themoviedb.org/reference/movie-account-states
type AccountState struct {
	ID        int         `json:"id"`
	Favorite  bool        `json:"favorite"`
	Rated     interface{} `json:"rated"` // Can be boolean (false) or RatedInfo object
	Watchlist bool        `json:"watchlist"`
}

// RatedInfo holds the rating value when an item is rated.
type RatedInfo struct {
	Value float64 `json:"value"`
}

// StatusResponse is a common response for POST/DELETE actions.
type StatusResponse struct {
	StatusCode    int    `json:"status_code"`
	StatusMessage string `json:"status_message"`
	Success       *bool  `json:"success,omitempty"` // Sometimes present
}

// AlternativeTitle represents an alternative title for a movie or TV show.
// See: https://developer.themoviedb.org/reference/movie-alternative-titles
type AlternativeTitle struct {
	ISO3166_1 string `json:"iso_3166_1"` // Country code
	Title     string `json:"title"`
	Type      string `json:"type"` // e.g., "", "Director's Cut"
}

// AlternativeTitlesResponse holds alternative titles.
type AlternativeTitlesResponse struct {
	ID      int                `json:"id"`                // Movie/TV ID
	Titles  []AlternativeTitle `json:"titles,omitempty"`  // Field name varies (titles/results)
	Results []AlternativeTitle `json:"results,omitempty"` // Field name varies (titles/results)
}

// Paginated represents common pagination fields.
type Paginated struct {
	Page         int `json:"page"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}
