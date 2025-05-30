package endpoints

import (
	"fmt"

	"github.com/falconer001/gotmdb/client"
	opts "github.com/falconer001/gotmdb/options"
	"github.com/falconer001/gotmdb/types"
)

// Movies handles communication with the movie related methods of the TMDb API.
// See: https://developer.themoviedb.org/reference/movie
type Movies struct {
	Client *client.Client
}

// GetDetails retrieves the primary information about a movie.
// Supports appending additional data like credits, images, videos, keywords,external_ids etc., using AppendToResponse.
// See: https://developer.themoviedb.org/reference/movie-details
func (m *Movies) GetDetails(movieID int) *opts.AppendToResponseBuilder[*types.MovieDetails] {
	return opts.NewAppendToResponseBuilder[*types.MovieDetails](m.Client, fmt.Sprintf("/movie/%d", movieID))
}

// GetRecommendations retrieves a list of recommended movies for a movie.
// See: https://developer.themoviedb.org/reference/movie-recommendations
func (m *Movies) GetRecommendations(movieID int) *opts.PagedBuilder[*types.MoviePaginatedResults] {
	return opts.NewPagedBuilder[*types.MoviePaginatedResults](m.Client, fmt.Sprintf("/movie/%d/recommendations", movieID))
}

// GetSimilar retrieves a list of similar movies.
// This is not the same as the "Recommendation" system you see on the website.
// These items are assembled by looking at keywords and genres.
// See: https://developer.themoviedb.org/reference/movie-similar
func (m *Movies) GetSimilar(movieID int) *opts.PagedBuilder[*types.MoviePaginatedResults] {
	return opts.NewPagedBuilder[*types.MoviePaginatedResults](m.Client, fmt.Sprintf("/movie/%d/recommendations", movieID))
}

// GetPopular retrieves a list of the current popular movies on TMDb.
// This list updates daily.
// See: https://developer.themoviedb.org/reference/movie-popular-list
func (m *Movies) GetPopular() *opts.PagedBuilder[*types.MoviePaginatedResults] {
	return opts.NewPagedBuilder[*types.MoviePaginatedResults](m.Client, "/movie/popular")
}

// GetTopRated retrieves the top rated movies on TMDb.
// See: https://developer.themoviedb.org/reference/movie-top-rated-list
func (m *Movies) GetTopRated() *opts.PagedBuilder[*types.MoviePaginatedResults] {
	return opts.NewPagedBuilder[*types.MoviePaginatedResults](m.Client, "/movie/top_rated")
}

// GetUpcoming retrieves a list of movies that are currently scheduled to be released.
// See: https://developer.themoviedb.org/reference/movie-upcoming-list
func (m *Movies) GetUpcoming() *opts.PagedBuilder[*types.UpcomingResponse] {
	return opts.NewPagedBuilder[*types.UpcomingResponse](m.Client, "/movie/upcoming")
}

// GetNowPlaying retrieves a list of movies that are currently playing in theatres.
// See: https://developer.themoviedb.org/reference/movie-now-playing-list
func (m *Movies) GetNowPlaying() *opts.PagedBuilder[*types.NowPlayingResponse] {
	return opts.NewPagedBuilder[*types.NowPlayingResponse](m.Client, "/movie/now_playing")
}

// GetLists retrieves a list of lists that this movie belongs to.
// See: https://developer.themoviedb.org/reference/movie-lists
func (m *Movies) GetLists(movieID int) *opts.PagedBuilder[*types.ListPaginatedResults] {
	return opts.NewPagedBuilder[*types.ListPaginatedResults](m.Client, fmt.Sprintf("/movie/%d/lists", movieID))
}

// GetReviews retrieves the user reviews for a movie.
// See: https://developer.themoviedb.org/reference/movie-reviews
func (m *Movies) GetReviews(movieID int) *opts.PagedBuilder[*types.ReviewPaginatedResults] {
	return opts.NewPagedBuilder[*types.ReviewPaginatedResults](m.Client, fmt.Sprintf("/movie/%d/reviews", movieID))
}

// GetAccountStates retrieves the rating, watchlist, and favorite status of a movie for a specific account.
// Requires either a SessionID or GuestSessionID (You will have to manually call the authservice to get them).
// See: https://developer.themoviedb.org/reference/movie-account-states
func (m *Movies) AccountStates(movieID int) *opts.StateSessionBuilder[*types.AccountState] {
	return opts.NewStateSessionBuilder[*types.AccountState](m.Client, fmt.Sprintf("/movie/%d/account_states", movieID), "GET", nil)
}

// RateMovie rates a movie.
// If not for guest session, bearer token is used (use .ForGuest(false)).
// A valid session or guest session ID is required (for guest session, use .ForGuest(true)).
// See: https://developer.themoviedb.org/reference/movie-add-rating
func (m *Movies) Rate(movieID int, body types.RatingRequest) *opts.StateSessionBuilder[*types.StatusResponse] {
	return opts.NewStateSessionBuilder[*types.StatusResponse](m.Client, fmt.Sprintf("/movie/%d/rating", movieID), "POST", body)
}

// DeleteRating removes your rating for a movie.
// If not for guest session, bearer token is used (use .ForGuest(false)).
// A valid session or guest session ID is required (for guest session, use .ForGuest(true)).
// See: https://developer.themoviedb.org/reference/movie-delete-rating
func (m *Movies) DeleteRating(movieID int) *opts.StateSessionBuilder[*types.StatusResponse] {
	return opts.NewStateSessionBuilder[*types.StatusResponse](m.Client, fmt.Sprintf("/movie/%d/rating", movieID), "DELETE", nil)
}

// GetChanges retrieves the changes for a movie.
// By default, only the last 24 hours of changes are returned.
// You can query up to 14 days in a single query by using the start_date and end_date opts.
// See: https://developer.themoviedb.org/reference/movie-changes
func (m *Movies) GetChanges(movieID int) *opts.ChangesBuilder {
	return opts.NewChangesBuilder(m.Client, fmt.Sprintf("/movie/%d/changes", movieID))
}

// GetKeywords retrieves the keywords that have been added to a movie.
// See: https://developer.themoviedb.org/reference/movie-keywords
func (m *Movies) GetKeywords(movieID int) *opts.NoOptsBuilder[*types.KeywordsResponse] {
	return opts.NewNoOptsBuilder[*types.KeywordsResponse](m.Client, fmt.Sprintf("/movie/%d/keywords", movieID))
}

// GetReleaseDates retrieves the release date and certification information for a movie.
// Returns release dates and certifications for a movie grouped by country.
// See: https://developer.themoviedb.org/reference/movie-release-dates
func (m *Movies) GetReleaseDates(movieID int) *opts.NoOptsBuilder[*types.ReleaseDatesResponse] {
	return opts.NewNoOptsBuilder[*types.ReleaseDatesResponse](m.Client, fmt.Sprintf("/movie/%d/release_dates", movieID))
}

// GetLatest retrieves the most recently created movie.
// Note: This endpoint might behave unexpectedly as it returns a single, potentially pre-release movie.
// See: https://developer.themoviedb.org/reference/movie-latestid
func (m *Movies) GetLatest() *opts.NoOptsBuilder[*types.MovieDetails] {
	return opts.NewNoOptsBuilder[*types.MovieDetails](m.Client, "/movie/latest")
}

// GetTranslations retrieves a list of translations that have been created for a movie.
// See: https://developer.themoviedb.org/reference/movie-translations
func (m *Movies) GetTranslations(movieID int) *opts.NoOptsBuilder[*types.TranslationsResponse] {
	return opts.NewNoOptsBuilder[*types.TranslationsResponse](m.Client, fmt.Sprintf("/movie/%d/translations", movieID))
}

// GetWatchProviders retrieves a list of the watch providers (streaming services, rental sources, etc.) for a movie.
// Powered by JustWatch.
// See: https://developer.themoviedb.org/reference/movie-watch-providers
func (m *Movies) GetWatchProviders(movieID int) *opts.NoOptsBuilder[*types.WatchProviderResponse] {
	return opts.NewNoOptsBuilder[*types.WatchProviderResponse](m.Client, fmt.Sprintf("/movie/%d/watch/providers", movieID))
}

// GetExternalIDs retrieves the external IDs for a movie (e.g., IMDb ID).
// See: https://developer.themoviedb.org/reference/movie-external-ids
func (m *Movies) GetExternalIDs(movieID int) *opts.NoOptsBuilder[*types.ExternalIDs] {
	return opts.NewNoOptsBuilder[*types.ExternalIDs](m.Client, fmt.Sprintf("/movie/%d/external_ids", movieID))
}

// GetAlternativeTitles retrieves the alternative titles for a movie.
// See: https://developer.themoviedb.org/reference/movie-alternative-titles
func (m *Movies) GetAlternativeTitles(movieID int) *opts.NoOptsBuilder[*types.AlternativeTitlesResponse] {
	return opts.NewNoOptsBuilder[*types.AlternativeTitlesResponse](m.Client, fmt.Sprintf("/movie/%d/alternative_titles", movieID))
}

// GetImages retrieves the images that belong to a movie.
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (like English) you can use the include_image_language parameter.
// See: https://developer.themoviedb.org/reference/movie-images
func (m *Movies) GetImages(movieID int) *opts.LangBuilder[*types.ImageList] {
	return opts.NewLangBuilder[*types.ImageList](m.Client, fmt.Sprintf("/movie/%d/images", movieID))
}

// GetVideos retrieves the videos that have been added to a movie (trailers, teasers, etc.).
// See: https://developer.themoviedb.org/reference/movie-videos
func (m *Movies) GetVideos(movieID int) *opts.LangBuilder[*types.VideoList] {
	return opts.NewLangBuilder[*types.VideoList](m.Client, fmt.Sprintf("/movie/%d/videos", movieID))
}

// GetCredits retrieves the cast and crew for a movie.
// See: https://developer.themoviedb.org/reference/movie-credits
func (m *Movies) GetCredits(movieID int) *opts.LangBuilder[*types.Credits] {
	return opts.NewLangBuilder[*types.Credits](m.Client, fmt.Sprintf("/movie/%d/credits", movieID))
}
