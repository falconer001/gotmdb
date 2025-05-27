package types

import "encoding/json"

// ListItem represents an item within a list (often a movie or TV show).
// The exact structure can vary, so using json.RawMessage might be necessary
// if a generic approach is needed, or define specific types if list type is known.
// For simplicity here, we assume it mirrors MovieListResult or TVListResult.
// Consider using a more robust approach like an interface or RawMessage if needed.
type ListItem json.RawMessage // Placeholder - Needs refinement based on actual usage

// ListDetails represents the details of a specific list.
// See: https://developer.themoviedb.org/reference/list-details
type ListDetails struct {
	CreatedBy     string     `json:"created_by"`
	Description   string     `json:"description"`
	FavoriteCount int        `json:"favorite_count"`
	ID            string     `json:"id"` // API docs say string, use string
	Items         []ListItem `json:"items"` // Array of movies or TV shows (structure varies)
	ItemCount     int        `json:"item_count"`
	ISO639_1      string     `json:"iso_639_1"` // Language code
	Name          string     `json:"name"`
	PosterPath    *string    `json:"poster_path"` // Nullable
}

// ListItemStatusResponse indicates if a specific movie is in a list.
// See: https://developer.themoviedb.org/reference/list-check-item-status
type ListItemStatusResponse struct {
	ID          string `json:"id"`           // The list ID
	ItemPresent bool   `json:"item_present"` // True if the movie_id is in the list
	// MediaID is not part of the response, it's a query param
}

// CreateListRequest is the request body for creating a new list.
// See: https://developer.themoviedb.org/reference/list-create
type CreateListRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Language    string `json:"language"` // ISO 639-1 language code
}

// CreateListResponse is the response after creating a list.
type CreateListResponse struct {
	StatusMessage string `json:"status_message"`
	Success       bool   `json:"success"`
	ListID        int    `json:"list_id"` // Note: The created list ID is an integer here
}

// AddItemRequest is the request body for adding an item to a list.
// See: https://developer.themoviedb.org/reference/list-add-movie
type AddItemRequest struct {
	MediaID int `json:"media_id"` // The ID of the movie or TV show to add
}

// RemoveItemRequest is the request body for removing an item from a list.
// See: https://developer.themoviedb.org/reference/list-remove-movie
type RemoveItemRequest struct {
	MediaID int `json:"media_id"` // The ID of the movie or TV show to remove
}

// List represents a list in paginated results (e.g., account lists).
// See: https://developer.themoviedb.org/reference/account-lists
type List struct {
	Description   string  `json:"description"`
	FavoriteCount int     `json:"favorite_count"`
	ID            int     `json:"id"` // Integer ID in this context
	ItemCount     int     `json:"item_count"`
	ISO639_1      string  `json:"iso_639_1"`
	ListType      string  `json:"list_type"` // e.g., "movie"
	Name          string  `json:"name"`
	PosterPath    *string `json:"poster_path"` // Nullable
}

// ListPaginatedResults represents paginated list results.
// Used for account lists and lists a movie belongs to.
type ListPaginatedResults struct {
	ID        *int   `json:"id,omitempty"` // Movie/Account ID (optional depending on context)
	Paginated        // Embed common pagination fields
	Results   []List `json:"results"`
}

