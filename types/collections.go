package types

// CollectionDetails represents the details of a movie collection.
// See: https://developer.themoviedb.org/reference/collection-details
type CollectionDetails struct {
	ID           int               `json:"id"`
	Name         string            `json:"name"`
	Overview     string            `json:"overview"`
	PosterPath   *string           `json:"poster_path"` // Nullable
	BackdropPath *string           `json:"backdrop_path"` // Nullable
	Parts        []MovieListResult `json:"parts"`        // List of movies in the collection
}

// CollectionImagesResponse holds the images for a collection.
// See: https://developer.themoviedb.org/reference/collection-images
type CollectionImagesResponse struct {
	ID        int     `json:"id"`
	Backdrops []Image `json:"backdrops"`
	Posters   []Image `json:"posters"`
}

// CollectionTranslationsResponse holds the translations for a collection.
// See: https://developer.themoviedb.org/reference/collection-translations
type CollectionTranslationsResponse struct {
	ID           int           `json:"id"`
	Translations []Translation `json:"translations"` // Uses the common Translation struct
}

// CollectionSearchResult represents a collection in search results.
// See: https://developer.themoviedb.org/reference/search-collection
type CollectionSearchResult struct {
	Adult            bool    `json:"adult"` // Included in API response but often false for collections
	BackdropPath     *string `json:"backdrop_path"` // Nullable
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	OriginalLanguage string  `json:"original_language"`
	OriginalName     string  `json:"original_name"`
	Overview         string  `json:"overview"`
	PosterPath       *string `json:"poster_path"` // Nullable
}

// CollectionSearchResponse represents paginated collection search results.
type CollectionSearchResponse struct {
	Paginated                 // Embed common pagination fields
	Results   []CollectionSearchResult `json:"results"`
}

