package types

// Certification represents a single content certification.
// Used within CertificationsResponse.
type Certification struct {
	Certification string `json:"certification"` // e.g., "PG-13", "TV-MA"
	Meaning       string `json:"meaning"`       // Description of the certification
	Order         int    `json:"order"`         // Sort order
}

// CertificationsResponse holds the certifications for a specific country.
// The map key is the country code (e.g., "US", "GB").
// See: https://developer.themoviedb.org/reference/certification-movie-list
type CertificationsResponse struct {
	Certifications map[string][]Certification `json:"certifications"`
}

