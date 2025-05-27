package types

// ChangeItem represents a single item in the general change lists.
// See: https://developer.themoviedb.org/reference/changes-movie-list
type ChangeItem struct {
	ID    int   `json:"id"`              // The movie, TV show, or person ID that has changed
	Adult *bool `json:"adult,omitempty"` // Optional adult status (primarily for movies)
}

// ChangeListResponse represents the paginated response for general change lists.
type ChangeListResponse struct {
	Paginated              // Embed common pagination fields
	Results   []ChangeItem `json:"results"`
}

// ChangeItemDetail represents a specific change made to an item.
// Used within ChangeGroup.
type ChangeItemDetail struct {
	ID            string  `json:"id"`                   // Internal ID of the change
	Action        string  `json:"action"`               // e.g., "added", "updated", "deleted"
	Time          string  `json:"time"`                 // Timestamp of the change
	ISO639_1      *string `json:"iso_639_1,omitempty"`  // Language code (if applicable)
	ISO3166_1     *string `json:"iso_3166_1,omitempty"` // Country code (if applicable)
	Value         any     `json:"value"`                // The new value (can be string, object, etc.)
	OriginalValue any     `json:"original_value"`       // The original value (can be string, object, etc.)
}

// ChangeGroup represents a group of changes for a specific key (field).
// Used in item-specific change endpoints.
type ChangeGroup struct {
	Key   string             `json:"key"`   // The field that changed (e.g., "overview", "images")
	Items []ChangeItemDetail `json:"items"` // List of changes for this key
}

// ItemChangesResponse represents the response for item-specific change endpoints (movie, tv, person, etc.).
// See: https://developer.themoviedb.org/reference/movie-changes
type ItemChangesResponse struct {
	Changes []ChangeGroup `json:"changes"`
}
