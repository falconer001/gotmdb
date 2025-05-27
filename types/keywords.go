package types

// KeywordSearchResponse represents paginated keyword search results.
// See: https://developer.themoviedb.org/reference/search-keyword
type KeywordSearchResponse struct {
	Paginated         // Embed common pagination fields
	Results   []Keyword `json:"results"` // Uses the common Keyword struct
}

// KeywordMoviesResponse represents paginated movie results for a specific keyword.
// Note: The API response includes the keyword ID at the top level.
// See: https://developer.themoviedb.org/reference/keyword-movies
type KeywordMoviesResponse struct {
	ID        int               `json:"id"` // The keyword ID
	Paginated                 // Embed common pagination fields
	Results   []MovieListResult `json:"results"` // Uses MovieListResult (defined in movies.go)
}

