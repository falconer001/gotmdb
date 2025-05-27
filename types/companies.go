package types

// ParentCompany represents a parent company, used within CompanyDetails.
type ParentCompany struct {
	ID       int     `json:"id"`
	LogoPath *string `json:"logo_path"` // Nullable
	Name     string  `json:"name"`
}

// CompanyDetails represents the details of a production company.
// See: https://developer.themoviedb.org/reference/company-details
type CompanyDetails struct {
	Description   string         `json:"description"`
	Headquarters  string         `json:"headquarters"`
	Homepage      string         `json:"homepage"`
	ID            int            `json:"id"`
	LogoPath      *string        `json:"logo_path"` // Nullable
	Name          string         `json:"name"`
	OriginCountry string         `json:"origin_country"`
	ParentCompany *ParentCompany `json:"parent_company,omitempty"` // Nullable
}

// AlternativeName represents an alternative name for a company or network.
type AlternativeName struct {
	Name string `json:"name"`
	Type string `json:"type"` // e.g., "Acronym", "Previous Name"
}

// CompanyAlternativeNamesResponse holds the alternative names for a company.
// See: https://developer.themoviedb.org/reference/company-alternative-names
type CompanyAlternativeNamesResponse struct {
	ID      int               `json:"id"`
	Results []AlternativeName `json:"results"`
}

// CompanyImagesResponse holds the logo images for a company.
// Note: Only logos are typically available for companies.
// See: https://developer.themoviedb.org/reference/company-images
type CompanyImagesResponse struct {
	ID    int     `json:"id"`
	Logos []Image `json:"logos"` // Uses the common Image struct
}

// CompanySearchResult represents a company in search results.
// See: https://developer.themoviedb.org/reference/search-company
type CompanySearchResult struct {
	ID            int     `json:"id"`
	LogoPath      *string `json:"logo_path"` // Nullable
	Name          string  `json:"name"`
	OriginCountry string  `json:"origin_country,omitempty"` // May not always be present
}

// CompanySearchResponse represents paginated company search results.
type CompanySearchResponse struct {
	Paginated               // Embed common pagination fields
	Results   []CompanySearchResult `json:"results"`
}

