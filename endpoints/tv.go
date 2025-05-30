package endpoints

import (
	"fmt"

	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/options"
	"github.com/falconer001/gotmdb/types"
)

type TV struct {
	Client *client.Client
}

// GetDetails retrieves the primary information about a TV series.
// Supports appending additional data like credits, images, videos, etc.
// See: https://developer.themoviedb.org/reference/tv-series-details
func (t *TV) GetDetails(seriesID int) *options.AppendToResponseBuilder[*types.TVDetails] {
	return options.NewAppendToResponseBuilder[*types.TVDetails](t.Client, fmt.Sprintf("/tv/%d", seriesID))
}

// GetRecommendations retrieves a list of recommended TV shows for a series.
// See: https://developer.themoviedb.org/reference/tv-series-recommendations
func (t *TV) GetRecommendations(seriesID int) *options.PagedBuilder[*types.TVShowPaginatedResults] {
	return options.NewPagedBuilder[*types.TVShowPaginatedResults](t.Client, fmt.Sprintf("/tv/%d/recommendations", seriesID))
}

// GetSimilar retrieves a list of similar TV shows.
// This is not the same as the "Recommendation" system you see on the website.
// These items are assembled by looking at keywords and genres.
// See: https://developer.themoviedb.org/reference/tv-series-similar
func (t *TV) GetSimilar(seriesID int) *options.PagedBuilder[*types.TVShowPaginatedResults] {
	return options.NewPagedBuilder[*types.TVShowPaginatedResults](t.Client, fmt.Sprintf("/tv/%d/similar", seriesID))
}

// GetPopular retrieves a list of the current popular TV shows on TMDb.
// This list updates daily.
// See: https://developer.themoviedb.org/reference/tv-series-popular
func (t *TV) GetPopular() *options.PagedBuilder[*types.TVShowPaginatedResults] {
	return options.NewPagedBuilder[*types.TVShowPaginatedResults](t.Client, "/tv/popular")
}

// GetOnTheAir retrieves a list of TV shows that are currently on the air.
// Note: This call is really just a discover call behind the scenes.
// See: https://developer.themoviedb.org/reference/tv-series-on-the-air
func (t *TV) GetOnTheAir() *options.PagedBuilder[*types.TVShowPaginatedResults] {
	return options.NewPagedBuilder[*types.TVShowPaginatedResults](t.Client, "/tv/on_the_air")
}

// GetAiringToday retrieves a list of TV shows that are airing today.
// See: https://developer.themoviedb.org/reference/tv-series-airing-today
func (t *TV) GetAiringToday() *options.PagedBuilder[*types.TVShowPaginatedResults] {
	return options.NewPagedBuilder[*types.TVShowPaginatedResults](t.Client, "/tv/airing_today")
}

// GetTopRated retrieves the top rated TV shows on TMDb.
// See: https://developer.themoviedb.org/reference/tv-series-top-rated
func (t *TV) GetTopRated() *options.PagedBuilder[*types.TVShowPaginatedResults] {
	return options.NewPagedBuilder[*types.TVShowPaginatedResults](t.Client, "/tv/top_rated")
}

// GetReviews retrieves the user reviews for a TV series.
// See: https://developer.themoviedb.org/reference/tv-series-reviews
func (t *TV) GetReviews(seriesID int) *options.PagedBuilder[*types.ReviewPaginatedResults] {
	return options.NewPagedBuilder[*types.ReviewPaginatedResults](t.Client, fmt.Sprintf("/tv/%d/reviews", seriesID))
}

// GetAccountStates retrieves the rating, watchlist, and favorite status of a TV show for a specific account.
// Requires either a SessionID or GuestSessionID.
// See: https://developer.themoviedb.org/reference/tv-series-account-states
func (t *TV) GetAccountStates(seriesID int) *options.StateSessionBuilder[*types.AccountState] {
	return options.NewStateSessionBuilder[*types.AccountState](t.Client, fmt.Sprintf("/tv/%d/account_states", seriesID), "GET", nil)
}

// Rate rates a TV series.
// A valid session or guest session ID is required.
// See: https://developer.themoviedb.org/reference/tv-series-add-rating
func (t *TV) Rate(seriesID int, body types.RatingRequest) *options.StateSessionBuilder[*types.StatusResponse] {
	return options.NewStateSessionBuilder[*types.StatusResponse](t.Client, fmt.Sprintf("/tv/%d/rating", seriesID), "POST", body)
}

// DeleteRating removes your rating for a TV series.
// A valid session or guest session ID is required.
// See: https://developer.themoviedb.org/reference/tv-series-delete-rating
func (t *TV) DeleteRating(seriesID int) *options.StateSessionBuilder[*types.StatusResponse] {
	return options.NewStateSessionBuilder[*types.StatusResponse](t.Client, fmt.Sprintf("/tv/%d/rating", seriesID), "DELETE", nil)
}

// GetAggregateCredits retrieves the aggregate cast and crew credits for a TV series.
// This call differs from the main credits call in that it does not return episode specific data.
// See: https://developer.themoviedb.org/reference/tv-series-aggregate-credits
func (t *TV) GetAggregateCredits(seriesID int) *options.LangBuilder[*types.Credits] {
	return options.NewLangBuilder[*types.Credits](t.Client, fmt.Sprintf("/tv/%d/aggregate_credits", seriesID))
}

// GetAlternativeTitles retrieves the alternative titles for a TV series.
// See: https://developer.themoviedb.org/reference/tv-series-alternative-titles
func (t *TV) GetAlternativeTitles(seriesID int) *options.LangBuilder[*types.AlternativeTitlesResponse] {
	return options.NewLangBuilder[*types.AlternativeTitlesResponse](t.Client, fmt.Sprintf("/tv/%d/alternative_titles", seriesID))
}

// GetContentRatings retrieves the content ratings for a TV series.
// See: https://developer.themoviedb.org/reference/tv-series-content-ratings
func (t *TV) GetContentRatings(seriesID int) *options.LangBuilder[*types.ContentRatingsResponse] {
	return options.NewLangBuilder[*types.ContentRatingsResponse](t.Client, fmt.Sprintf("/tv/%d/content_ratings", seriesID))
}

// GetCredits retrieves the credits (cast, crew, guest stars) for a TV series.
// Note: This is different from AggregateCredits as it includes episode-level data.
// See: https://developer.themoviedb.org/reference/tv-series-credits
func (t *TV) GetCredits(seriesID int) *options.LangBuilder[*types.Credits] {
	return options.NewLangBuilder[*types.Credits](t.Client, fmt.Sprintf("/tv/%d/credits", seriesID))
}

// GetEpisodeGroups retrieves the episode groups that have been created for a TV series.
// See: https://developer.themoviedb.org/reference/tv-series-episode-groups
func (t *TV) GetEpisodeGroups(seriesID int) *options.LangBuilder[*types.EpisodeGroupsResponse] {
	return options.NewLangBuilder[*types.EpisodeGroupsResponse](t.Client, fmt.Sprintf("/tv/%d/episode_groups", seriesID))
}

// GetExternalIDs retrieves the external IDs for a TV series (e.g., IMDb ID, TVDB ID).
// See: https://developer.themoviedb.org/reference/tv-series-external-ids
func (t *TV) GetExternalIDs(seriesID int) *options.LangBuilder[*types.TVExternalIDs] {
	return options.NewLangBuilder[*types.TVExternalIDs](t.Client, fmt.Sprintf("/tv/%d/external_ids", seriesID))
}

// GetKeywords retrieves the keywords that have been added to a TV series.
// See: https://developer.themoviedb.org/reference/tv-series-keywords
func (t *TV) GetKeywords(seriesID int) *options.LangBuilder[*types.KeywordsResponse] {
	return options.NewLangBuilder[*types.KeywordsResponse](t.Client, fmt.Sprintf("/tv/%d/keywords", seriesID))
}

// GetImages retrieves the images that belong to a TV series.
// See: https://developer.themoviedb.org/reference/tv-series-images
func (t *TV) GetImages(seriesID int) *options.LangBuilder[*types.ImageList] {
	return options.NewLangBuilder[*types.ImageList](t.Client, fmt.Sprintf("/tv/%d/images", seriesID))
}

// GetVideos retrieves the videos that have been added to a TV series (trailers, teasers, etc.).
// See: https://developer.themoviedb.org/reference/tv-series-videos
func (t *TV) GetVideos(seriesID int) *options.LangBuilder[*types.VideoList] {
	return options.NewLangBuilder[*types.VideoList](t.Client, fmt.Sprintf("/tv/%d/videos", seriesID))
}

// GetScreenedTheatrically retrieves a list of seasons or episodes that have been screened theatrically.
// See: https://developer.themoviedb.org/reference/tv-series-screened-theatrically
func (t *TV) GetScreenedTheatrically(seriesID int) *options.NoOptsBuilder[*types.ScreenedTheatricallyResponse] {
	return options.NewNoOptsBuilder[*types.ScreenedTheatricallyResponse](t.Client, fmt.Sprintf("/tv/%d/screened_theatrically", seriesID))
}

// GetTranslations retrieves a list of translations that have been created for a TV series.
// See: https://developer.themoviedb.org/reference/tv-series-translations
func (t *TV) GetTranslations(seriesID int) *options.NoOptsBuilder[*types.TranslationsResponse] {
	return options.NewNoOptsBuilder[*types.TranslationsResponse](t.Client, fmt.Sprintf("/tv/%d/translations", seriesID))
}

// GetWatchProviders retrieves a list of the watch providers (streaming services, rental sources, etc.) for a TV series.
// Powered by JustWatch.
// See: https://developer.themoviedb.org/reference/tv-series-watch-providers
func (t *TV) GetWatchProviders(seriesID int) *options.NoOptsBuilder[*types.WatchProviderResponse] {
	return options.NewNoOptsBuilder[*types.WatchProviderResponse](t.Client, fmt.Sprintf("/tv/%d/watch/providers", seriesID))
}

// GetChanges retrieves the changes for a TV series.
// By default, only the last 24 hours of changes are returned.
// You can query up to 14 days in a single query by using the start_date and end_date options.
// See: https://developer.themoviedb.org/reference/tv-series-changes
func (t *TV) GetChanges(seriesID int) *options.ChangesBuilder {
	return options.NewChangesBuilder(t.Client, fmt.Sprintf("/tv/%d/changes", seriesID))
}
