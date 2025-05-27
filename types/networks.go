package types

// Network represents a TV network.
// See: https://developer.themoviedb.org/reference/network-details
type Network struct {
	Headquarters  string  `json:"headquarters"`
	Homepage      string  `json:"homepage"`
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"` // Nullable
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country"`
}

// NetworkAlternativeNamesResponse holds the alternative names for a network.
// See: https://developer.themoviedb.org/reference/network-alternative-names
type NetworkAlternativeNamesResponse struct {
	ID      int               `json:"id"`
	Results []AlternativeName `json:"results"` // Uses common AlternativeName
}

// NetworkImagesResponse holds the logo images for a network.
// See: https://developer.themoviedb.org/reference/network-images
type NetworkImagesResponse struct {
	ID    int     `json:"id"`
	Logos []Image `json:"logos"` // Uses common Image
}

