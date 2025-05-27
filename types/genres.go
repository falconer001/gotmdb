package types

// GenreListResponse holds a list of genres.
// See: https://developer.themoviedb.org/reference/genre-movie-list
type GenreListResponse struct {
	Genres []Genre `json:"genres"` // Uses the common Genre struct
}

