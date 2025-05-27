package types

// ReviewDetails represents the detailed information for a specific review.
// It extends the common Review struct with media information.
// See: https://developer.themoviedb.org/reference/review-details
type ReviewDetails struct {
	ID            string        `json:"id"` // Review ID
	Author        string        `json:"author"`
	AuthorDetails AuthorDetails `json:"author_details"` // Uses common AuthorDetails
	Content       string        `json:"content"`
	CreatedAt     string        `json:"created_at"` // Timestamp
	ISO639_1      string        `json:"iso_639_1"`    // Language code of the review
	MediaID       int           `json:"media_id"`     // ID of the movie or TV show reviewed
	MediaTitle    string        `json:"media_title"`  // Title of the movie or TV show reviewed
	MediaType     string        `json:"media_type"`   // "movie" or "tv"
	UpdatedAt     string        `json:"updated_at"` // Timestamp
	URL           string        `json:"url"`
}

