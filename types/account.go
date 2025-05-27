package types

// AccountDetails represents the details of a TMDB user account.
// See: https://developer.themoviedb.org/reference/account-details
type AccountDetails struct {
	Avatar struct {
		Gravatar struct {
			Hash string `json:"hash"`
		} `json:"gravatar"`
		Tmdb struct {
			AvatarPath *string `json:"avatar_path"` // Nullable
		} `json:"tmdb"`
	} `json:"avatar"`
	ID            int    `json:"id"`
	ISO639_1      string `json:"iso_639_1"`      // Language code
	ISO3166_1     string `json:"iso_3166_1"`     // Country code
	Name          string `json:"name"`          // Full name or display name
	IncludeAdult bool   `json:"include_adult"` // Whether to include adult content
	Username      string `json:"username"`
}

// MarkFavoriteRequest is the request body for marking an item as favorite.
// See: https://developer.themoviedb.org/reference/account-add-favorite
type MarkFavoriteRequest struct {
	MediaType string `json:"media_type"` // "movie" or "tv"
	MediaID   int    `json:"media_id"`   // The ID of the movie or TV show
	Favorite  bool   `json:"favorite"`  // true to mark as favorite, false to unmark
}

// AddToWatchlistRequest is the request body for adding an item to the watchlist.
// See: https://developer.themoviedb.org/reference/account-add-to-watchlist
type AddToWatchlistRequest struct {
	MediaType string `json:"media_type"` // "movie" or "tv"
	MediaID   int    `json:"media_id"`   // The ID of the movie or TV show
	Watchlist bool   `json:"watchlist"`  // true to add to watchlist, false to remove
}

// RatedMovie represents a movie entry in the rated movies list.
// It extends MovieListResult with a 'rating' field.
type RatedMovie struct {
	MovieListResult           // Embed basic movie list info
	Rating          float64 `json:"rating"` // User's rating for this movie
}

// RatedMoviePaginatedResults represents paginated results for rated movies.
type RatedMoviePaginatedResults struct {
	Paginated             // Embed common pagination fields
	Results   []RatedMovie `json:"results"`
}

// RatedTVShow represents a TV show entry in the rated TV shows list.
// It extends TVListResult with a 'rating' field.
type RatedTVShow struct {
	TVListResult           // Embed basic TV list info
	Rating       float64 `json:"rating"` // User's rating for this TV show
}

// RatedTVShowPaginatedResults represents paginated results for rated TV shows.
type RatedTVShowPaginatedResults struct {
	Paginated             // Embed common pagination fields
	Results   []RatedTVShow `json:"results"`
}

// RatedTVEpisode represents a TV episode entry in the rated TV episodes list.
// It extends TVEpisodeListResult (which needs definition) with a 'rating' field.
type RatedTVEpisode struct {
	TVEpisodeListResult           // Embed basic TV episode list info
	Rating            float64 `json:"rating"` // User's rating for this episode
}

// RatedTVEpisodePaginatedResults represents paginated results for rated TV episodes.
type RatedTVEpisodePaginatedResults struct {
	Paginated                // Embed common pagination fields
	Results   []RatedTVEpisode `json:"results"`
}

