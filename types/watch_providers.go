package types

// WatchProviderRegion represents a region where watch providers are available.
// See: https://developer.themoviedb.org/reference/watch-providers-available-regions
type WatchProviderRegion struct {
	ISO3166_1   string `json:"iso_3166_1"`   // e.g., "US"
	EnglishName string `json:"english_name"` // e.g., "United States"
	NativeName  string `json:"native_name"`  // e.g., "United States"
}

// WatchProviderRegionsResponse holds the list of available watch provider regions.
type WatchProviderRegionsResponse struct {
	Results []WatchProviderRegion `json:"results"`
}

// WatchProviderInfo represents basic information about a watch provider.
// Used in the general provider lists.
// See: https://developer.themoviedb.org/reference/watch-providers-list
type WatchProviderInfo struct {
	DisplayPriorities map[string]int `json:"display_priorities,omitempty"` // Map of country code to priority
	DisplayPriority   int            `json:"display_priority"`           // Default priority
	LogoPath          string         `json:"logo_path"`
	ProviderName      string         `json:"provider_name"`
	ProviderID        int            `json:"provider_id"`
}

// WatchProviderListResponse holds the list of available watch providers.
type WatchProviderListResponse struct {
	Results []WatchProviderInfo `json:"results"`
}

// WatchProvider represents a specific provider offering a movie/TV show.
// Used within CountryWatchProviders.
type WatchProvider struct {
	LogoPath        string `json:"logo_path"`
	ProviderID      int    `json:"provider_id"`
	ProviderName    string `json:"provider_name"`
	DisplayPriority int    `json:"display_priority"`
}

// CountryWatchProviders holds the different ways (flatrate, rent, buy, etc.)
// a movie/TV show is available from providers in a specific country.
type CountryWatchProviders struct {
	Link     string          `json:"link"` // Link to the TMDB watch page for this country
	Flatrate []WatchProvider `json:"flatrate,omitempty"`
	Rent     []WatchProvider `json:"rent,omitempty"`
	Buy      []WatchProvider `json:"buy,omitempty"`
	Free     []WatchProvider `json:"free,omitempty"`
	Ads      []WatchProvider `json:"ads,omitempty"`
}

// WatchProviderResponse holds the watch provider information for a movie or TV show,
// grouped by country code (e.g., "US", "GB").
// See: https://developer.themoviedb.org/reference/movie-watch-providers
type WatchProviderResponse struct {
	ID      int                              `json:"id"` // Movie or TV Show ID
	Results map[string]CountryWatchProviders `json:"results"`
}

