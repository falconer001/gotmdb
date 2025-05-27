package types

// PersonDetails represents the detailed information for a person.
// Includes fields for append_to_response results when requested.
// See: https://developer.themoviedb.org/reference/person-details
type PersonDetails struct {
	Adult              bool      `json:"adult"`
	AlsoKnownAs        []string  `json:"also_known_as"`
	Biography          string    `json:"biography"`
	Birthday           *string   `json:"birthday"`           // Nullable, Format YYYY-MM-DD
	Deathday           *string   `json:"deathday"`           // Nullable, Format YYYY-MM-DD
	Gender             int       `json:"gender"`             // 0: Not set, 1: Female, 2: Male, 3: Non-binary
	Homepage           *string   `json:"homepage"`           // Nullable
	ID                 int       `json:"id"`
	IMDbID             *string   `json:"imdb_id"`            // Nullable, Format nm1234567
	KnownForDepartment string    `json:"known_for_department"` // e.g., "Acting"
	Name               string    `json:"name"`
	PlaceOfBirth       *string   `json:"place_of_birth"`      // Nullable
	Popularity         float64   `json:"popularity"`
	ProfilePath        *string   `json:"profile_path"`       // Nullable

	// --- Appended fields --- 
	MovieCredits    *PersonMovieCreditsResponse    `json:"movie_credits,omitempty"`
	TVCredits       *PersonTVCreditsResponse       `json:"tv_credits,omitempty"`
	CombinedCredits *PersonCombinedCreditsResponse `json:"combined_credits,omitempty"`
	ExternalIDs     *PersonExternalIDs             `json:"external_ids,omitempty"`
	Images          *PersonImagesResponse          `json:"images,omitempty"`
	TaggedImages    *TaggedImagePaginatedResponse  `json:"tagged_images,omitempty"`
	Translations    *PersonTranslationsResponse    `json:"translations,omitempty"`
}

// PersonMovieCastCredit represents a movie cast credit for a person.
// Extends MovieListResult with character and credit_id.
type PersonMovieCastCredit struct {
	MovieListResult           // Embed basic movie list info
	Character       string `json:"character"`
	CreditID        string `json:"credit_id"`
}

// PersonMovieCrewCredit represents a movie crew credit for a person.
// Extends MovieListResult with department, job, and credit_id.
type PersonMovieCrewCredit struct {
	MovieListResult           // Embed basic movie list info
	Department      string `json:"department"`
	Job             string `json:"job"`
	CreditID        string `json:"credit_id"`
}

// PersonMovieCreditsResponse holds a person's movie credits.
// See: https://developer.themoviedb.org/reference/person-movie-credits
type PersonMovieCreditsResponse struct {
	ID   int                     `json:"id"` // Person ID
	Cast []PersonMovieCastCredit `json:"cast"`
	Crew []PersonMovieCrewCredit `json:"crew"`
}

// PersonTVCastCredit represents a TV cast credit for a person.
// Extends TVListResult with character, credit_id, and episode_count.
type PersonTVCastCredit struct {
	TVListResult              // Embed basic TV list info
	Character    string `json:"character"`
	CreditID     string `json:"credit_id"`
	EpisodeCount int    `json:"episode_count"`
}

// PersonTVCrewCredit represents a TV crew credit for a person.
// Extends TVListResult with department, job, credit_id, and episode_count.
type PersonTVCrewCredit struct {
	TVListResult              // Embed basic TV list info
	Department   string `json:"department"`
	Job          string `json:"job"`
	CreditID     string `json:"credit_id"`
	EpisodeCount int    `json:"episode_count"`
}

// PersonTVCreditsResponse holds a person's TV credits.
// See: https://developer.themoviedb.org/reference/person-tv-credits
type PersonTVCreditsResponse struct {
	ID   int                  `json:"id"` // Person ID
	Cast []PersonTVCastCredit `json:"cast"`
	Crew []PersonTVCrewCredit `json:"crew"`
}

// PersonCombinedCastCredit represents a combined movie or TV cast credit.
// Includes media_type to distinguish.
type PersonCombinedCastCredit struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"` // Nullable
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"` // Movie or TV Show ID
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    *string  `json:"original_title,omitempty"` // Movie
	OriginalName     *string  `json:"original_name,omitempty"`  // TV
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"` // Nullable
	ReleaseDate      *string  `json:"release_date,omitempty"` // Movie (YYYY-MM-DD)
	FirstAirDate     *string  `json:"first_air_date,omitempty"` // TV (YYYY-MM-DD)
	Title            *string  `json:"title,omitempty"` // Movie
	Name             *string  `json:"name,omitempty"`  // TV
	Video            *bool    `json:"video,omitempty"` // Movie
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Character        string   `json:"character"`
	CreditID         string   `json:"credit_id"`
	Order            *int     `json:"order,omitempty"` // Movie cast order
	MediaType        string   `json:"media_type"` // "movie" or "tv"
	EpisodeCount     *int     `json:"episode_count,omitempty"` // TV
	OriginCountry    []string `json:"origin_country,omitempty"` // TV
}

// PersonCombinedCrewCredit represents a combined movie or TV crew credit.
// Includes media_type to distinguish.
type PersonCombinedCrewCredit struct {
	Adult            bool     `json:"adult"`
	BackdropPath     *string  `json:"backdrop_path"` // Nullable
	GenreIDs         []int    `json:"genre_ids"`
	ID               int      `json:"id"` // Movie or TV Show ID
	OriginalLanguage string   `json:"original_language"`
	OriginalTitle    *string  `json:"original_title,omitempty"` // Movie
	OriginalName     *string  `json:"original_name,omitempty"`  // TV
	Overview         string   `json:"overview"`
	Popularity       float64  `json:"popularity"`
	PosterPath       *string  `json:"poster_path"` // Nullable
	ReleaseDate      *string  `json:"release_date,omitempty"` // Movie (YYYY-MM-DD)
	FirstAirDate     *string  `json:"first_air_date,omitempty"` // TV (YYYY-MM-DD)
	Title            *string  `json:"title,omitempty"` // Movie
	Name             *string  `json:"name,omitempty"`  // TV
	Video            *bool    `json:"video,omitempty"` // Movie
	VoteAverage      float64  `json:"vote_average"`
	VoteCount        int      `json:"vote_count"`
	Department       string   `json:"department"`
	Job              string   `json:"job"`
	CreditID         string   `json:"credit_id"`
	MediaType        string   `json:"media_type"` // "movie" or "tv"
	EpisodeCount     *int     `json:"episode_count,omitempty"` // TV
	OriginCountry    []string `json:"origin_country,omitempty"` // TV
}

// PersonCombinedCreditsResponse holds a person's combined movie and TV credits.
// See: https://developer.themoviedb.org/reference/person-combined-credits
type PersonCombinedCreditsResponse struct {
	ID   int                        `json:"id"` // Person ID
	Cast []PersonCombinedCastCredit `json:"cast"`
	Crew []PersonCombinedCrewCredit `json:"crew"`
}

// PersonExternalIDs represents external IDs specifically for a person.
// See: https://developer.themoviedb.org/reference/person-external-ids
type PersonExternalIDs struct {
	ID          int     `json:"id"` // Person ID
	IMDbID      *string `json:"imdb_id,omitempty"`
	WikidataID  *string `json:"wikidata_id,omitempty"`
	FacebookID  *string `json:"facebook_id,omitempty"`
	InstagramID *string `json:"instagram_id,omitempty"`
	TwitterID   *string `json:"twitter_id,omitempty"`
	FreebaseMID *string `json:"freebase_mid,omitempty"` // Deprecated?
	FreebaseID  *string `json:"freebase_id,omitempty"`  // Deprecated?
	TVRageID    *int    `json:"tvrage_id,omitempty"`    // Deprecated?
}

// PersonImagesResponse holds profile images for a person.
// See: https://developer.themoviedb.org/reference/person-images
type PersonImagesResponse struct {
	ID       int     `json:"id"` // Person ID
	Profiles []Image `json:"profiles"` // Uses common Image struct
}

// MediaInfoShort represents abbreviated media info used in TaggedImage.
type MediaInfoShort struct {
	ID           int     `json:"id"`
	Name         *string `json:"name,omitempty"` // TV
	Title        *string `json:"title,omitempty"` // Movie
	OriginalName *string `json:"original_name,omitempty"` // TV
	OriginalTitle *string `json:"original_title,omitempty"` // Movie
	PosterPath   *string `json:"poster_path"`
	Adult        *bool   `json:"adult,omitempty"` // Movie
	Overview     string  `json:"overview"`
	ReleaseDate  *string `json:"release_date,omitempty"` // Movie
	FirstAirDate *string `json:"first_air_date,omitempty"` // TV
	VoteAverage  float64 `json:"vote_average"`
	VoteCount    int     `json:"vote_count"`
	MediaType    string  `json:"media_type"` // Added for context
}

// TaggedImage represents an image tagged with a person.
// See: https://developer.themoviedb.org/reference/person-tagged-images
type TaggedImage struct {
	AspectRatio float64        `json:"aspect_ratio"`
	FilePath    string         `json:"file_path"`
	Height      int            `json:"height"`
	ID          string         `json:"id"` // Image ID (string)
	ISO639_1    *string        `json:"iso_639_1"` // Nullable
	VoteAverage float64        `json:"vote_average"`
	VoteCount   int            `json:"vote_count"`
	Width       int            `json:"width"`
	ImageType   string         `json:"image_type"` // e.g., "poster", "backdrop"
	Media       MediaInfoShort `json:"media"`
	MediaType   string         `json:"media_type"` // "movie" or "tv"
}

// TaggedImagePaginatedResponse represents paginated tagged image results.
type TaggedImagePaginatedResponse struct {
	ID        int           `json:"id"` // Person ID
	Paginated               // Embed common pagination fields
	Results   []TaggedImage `json:"results"`
}

// PersonTranslationData holds the translatable biography field for a person.
type PersonTranslationData struct {
	Biography string `json:"biography"`
}

// PersonTranslation represents a translation for a person.
type PersonTranslation struct {
	ISO3166_1   string                `json:"iso_3166_1"`
	ISO639_1    string                `json:"iso_639_1"`
	Name        string                `json:"name"`
	EnglishName string                `json:"english_name"`
	Data        PersonTranslationData `json:"data"`
}

// PersonTranslationsResponse holds translations for a person.
// See: https://developer.themoviedb.org/reference/person-translations
type PersonTranslationsResponse struct {
	ID           int                 `json:"id"` // Person ID
	Translations []PersonTranslation `json:"translations"`
}

// KnownForItem represents a movie or TV show a person is known for.
// Used in PersonListResult.
type KnownForItem struct {
	Adult            bool    `json:"adult"` // Movie specific?
	BackdropPath     *string `json:"backdrop_path"`
	ID               int     `json:"id"` // Movie or TV ID
	Title            *string `json:"title,omitempty"` // Movie
	Name             *string `json:"name,omitempty"` // TV
	OriginalLanguage string  `json:"original_language"`
	OriginalTitle    *string `json:"original_title,omitempty"` // Movie
	OriginalName     *string `json:"original_name,omitempty"` // TV
	Overview         string  `json:"overview"`
	PosterPath       *string `json:"poster_path"`
	MediaType        string  `json:"media_type"` // "movie" or "tv"
	ReleaseDate      *string `json:"release_date,omitempty"` // Movie
	FirstAirDate     *string `json:"first_air_date,omitempty"` // TV
	Video            *bool   `json:"video,omitempty"` // Movie
	VoteAverage      float64 `json:"vote_average"`
	VoteCount        int     `json:"vote_count"`
	OriginCountry    []string `json:"origin_country,omitempty"` // TV
	GenreIDs         []int   `json:"genre_ids"`
	Popularity       float64 `json:"popularity"`
}

// PersonListResult represents a person in list results (e.g., popular people).
// See: https://developer.themoviedb.org/reference/person-popular-list
type PersonListResult struct {
	Adult              bool           `json:"adult"`
	Gender             int            `json:"gender"`
	ID                 int            `json:"id"`
	KnownForDepartment string         `json:"known_for_department"`
	KnownFor           []KnownForItem `json:"known_for"`
	Name               string         `json:"name"`
	OriginalName       string         `json:"original_name"`
	Popularity         float64        `json:"popularity"`
	ProfilePath        *string        `json:"profile_path"`
}

// PersonPaginatedResults represents paginated person list results.
type PersonPaginatedResults struct {
	Paginated         // Embed common pagination fields
	Results   []PersonListResult `json:"results"`
}

