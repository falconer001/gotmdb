package types

// ImageConfiguration holds the base URL and available sizes for images.
// Part of APIConfiguration.
type ImageConfiguration struct {
	BaseURL       string   `json:"base_url"`        // e.g., "http://image.tmdb.org/t/p/"
	SecureBaseURL string   `json:"secure_base_url"` // e.g., "https://image.tmdb.org/t/p/"
	BackdropSizes []string `json:"backdrop_sizes"` // e.g., ["w300", "w780", "w1280", "original"]
	LogoSizes     []string `json:"logo_sizes"`     // e.g., ["w45", "w92", "w154", "w185", "w300", "w500", "original"]
	PosterSizes   []string `json:"poster_sizes"`   // e.g., ["w92", "w154", "w185", "w342", "w500", "w780", "original"]
	ProfileSizes  []string `json:"profile_sizes"`  // e.g., ["w45", "w185", "h632", "original"]
	StillSizes    []string `json:"still_sizes"`    // e.g., ["w92", "w185", "w300", "original"]
}

// APIConfiguration represents the response from the main configuration endpoint.
// See: https://developer.themoviedb.org/reference/configuration-api
type APIConfiguration struct {
	Images      ImageConfiguration `json:"images"`
	ChangeKeys []string           `json:"change_keys"` // List of valid change keys
}

// Country represents a country supported by TMDB.
// See: https://developer.themoviedb.org/reference/configuration-countries
type Country struct {
	ISO3166_1   string `json:"iso_3166_1"`             // e.g., "US"
	EnglishName string `json:"english_name"`           // e.g., "United States of America"
	NativeName  string `json:"native_name,omitempty"` // e.g., "United States", may be empty
}

// JobDepartment represents a department and the jobs within it.
// See: https://developer.themoviedb.org/reference/configuration-jobs
type JobDepartment struct {
	Department string   `json:"department"` // e.g., "Directing"
	Jobs       []string `json:"jobs"`       // e.g., ["Director", "Assistant Director"]
}

// LanguageConfig represents a language supported by TMDB.
// See: https://developer.themoviedb.org/reference/configuration-languages
type LanguageConfig struct {
	ISO639_1    string `json:"iso_639_1"`    // e.g., "en"
	EnglishName string `json:"english_name"` // e.g., "English"
	Name        string `json:"name"`         // e.g., "English"
}

// Timezone represents a timezone supported by TMDB.
// See: https://developer.themoviedb.org/reference/configuration-timezones
type Timezone struct {
	ISO3166_1 string   `json:"iso_3166_1"` // Country code
	Zones     []string `json:"zones"`      // List of timezone strings (e.g., "America/New_York")
}

