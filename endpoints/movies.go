package endpoints

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/falconer001/gotmdb/client"
	"github.com/falconer001/gotmdb/types"
	"github.com/falconer001/gotmdb/utils"
)

// MoviesService handles communication with the movie related
// methods of the TMDb API.
// See: https://developer.themoviedb.org/reference/movie
type MoviesService struct {
	Client *client.Client
}

type LanguagePageOptions struct {
	Language string `url:"language,omitempty"`
	Page     int    `url:"page,omitempty"`
}

type LanguageOptions struct {
	Language string `url:"language,omitempty"`
}

type AuthFuncOptions struct {
	ForGuest       bool
	SessionID      *string `url:"session_id,omitempty"`
	GuestSessionID *string `url:"guest_session_id,omitempty"`
}

// MovieDetailsOptions represents the parameters for the GetDetails endpoint.
type MovieDetailsOptions struct {
	Language         *string  `url:"language,omitempty"`
	AppendToResponse []string `url:"-"` // Handled manually in the path
}

// GetDetails retrieves the primary information about a movie.
// Supports appending additional data like credits, images, videos, keywords,external_ids etc., using AppendToResponse.
// See: https://developer.themoviedb.org/reference/movie-details
func (ms *MoviesService) GetDetails(movieID int, opts *MovieDetailsOptions) (*types.MovieDetails, error) {
	path := fmt.Sprintf("/movie/%d", movieID)
	resp := new(types.MovieDetails)

	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	} else {
		params = url.Values{}
	}

	if opts != nil && len(opts.AppendToResponse) > 0 {
		params.Set("append_to_response", strings.Join(opts.AppendToResponse, ","))
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetAccountStates retrieves the rating, watchlist, and favorite status of a movie for a specific account.
// If ForGuest is true, either SessionID or GuestSessionID is required.
// Requires either a SessionID or GuestSessionID (You will have to manually call the authservice to get them).
// See: https://developer.themoviedb.org/reference/movie-account-states
func (ms *MoviesService) GetAccountStates(movieID int, opts *AuthFuncOptions) (*types.AccountState, error) {
	if opts != nil && opts.ForGuest {
		if opts.SessionID == nil && opts.GuestSessionID == nil {
			return nil, fmt.Errorf("either SessionID or GuestSessionID is required for GetAccountStates")
		}
	}

	path := fmt.Sprintf("/movie/%d/account_states", movieID)
	resp := new(types.AccountState)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// AlternativeTitlesOptions represents the available options for the GetAlternativeTitles endpoint.
type AlternativeTitlesOptions struct {
	Country *string `url:"country,omitempty"` // ISO 3166-1 code
}

// GetAlternativeTitles retrieves the alternative titles for a movie.
// See: https://developer.themoviedb.org/reference/movie-alternative-titles
func (ms *MoviesService) GetAlternativeTitles(movieID int, opts *AlternativeTitlesOptions) (*types.AlternativeTitlesResponse, error) {
	path := fmt.Sprintf("/movie/%d/alternative_titles", movieID)
	resp := new(types.AlternativeTitlesResponse)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// MovieChangesOptions represents the available options for the GetChanges endpoint.
type MovieChangesOptions struct {
	StartDate *string `url:"start_date,omitempty"` // Format YYYY-MM-DD
	EndDate   *string `url:"end_date,omitempty"`   // Format YYYY-MM-DD
	Page      *int    `url:"page,omitempty"`
}

// GetChanges retrieves the changes for a movie.
// By default, only the last 24 hours of changes are returned.
// You can query up to 14 days in a single query by using the start_date and end_date options.
// See: https://developer.themoviedb.org/reference/movie-changes
func (ms *MoviesService) GetChanges(movieID int, opts *MovieChangesOptions) (*types.ItemChangesResponse, error) {
	path := fmt.Sprintf("/movie/%d/changes", movieID)
	resp := new(types.ItemChangesResponse)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetCredits retrieves the cast and crew for a movie.
// See: https://developer.themoviedb.org/reference/movie-credits
func (ms *MoviesService) GetCredits(movieID int, opts *LanguageOptions) (*types.Credits, error) {
	path := fmt.Sprintf("/movie/%d/credits", movieID)
	resp := new(types.Credits)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// GetExternalIDs retrieves the external IDs for a movie (e.g., IMDb ID).
// See: https://developer.themoviedb.org/reference/movie-external-ids
func (ms *MoviesService) GetExternalIDs(movieID int) (*types.ExternalIDs, error) {
	path := fmt.Sprintf("/movie/%d/external_ids", movieID)
	resp := new(types.ExternalIDs)

	err := ms.Client.DoRequest("GET", path, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// MovieImagesOptions represents the available options for the GetImages endpoint.
type MovieImagesOptions struct {
	IncludeImageLanguage *string `url:"include_image_language,omitempty"` // Comma separated ISO 639-1 codes
	Language             *string `url:"language,omitempty"`
}

// GetImages retrieves the images that belong to a movie.
// Querying images with a language parameter will filter the results.
// If you want to include a fallback language (like English) you can use the include_image_language parameter.
// See: https://developer.themoviedb.org/reference/movie-images
func (ms *MoviesService) GetImages(movieID int, opts *MovieImagesOptions) (*types.ImageList, error) {
	path := fmt.Sprintf("/movie/%d/images", movieID)
	resp := new(types.ImageList)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// GetKeywords retrieves the keywords that have been added to a movie.
// See: https://developer.themoviedb.org/reference/movie-keywords
func (ms *MoviesService) GetKeywords(movieID int) (*types.KeywordsResponse, error) {
	path := fmt.Sprintf("/movie/%d/keywords", movieID)
	resp := new(types.KeywordsResponse)

	err := ms.Client.DoRequest("GET", path, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// GetLists retrieves a list of lists that this movie belongs to.
// See: https://developer.themoviedb.org/reference/movie-lists
func (ms *MoviesService) GetLists(movieID int, opts *LanguagePageOptions) (*types.ListPaginatedResults, error) {
	path := fmt.Sprintf("/movie/%d/lists", movieID)
	resp := new(types.ListPaginatedResults)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = &movieID
	return resp, nil
}

// GetRecommendations retrieves a list of recommended movies for a movie.
// See: https://developer.themoviedb.org/reference/movie-recommendations
func (ms *MoviesService) GetRecommendations(movieID int, opts *LanguagePageOptions) (*types.MoviePaginatedResults, error) {
	path := fmt.Sprintf("/movie/%d/recommendations", movieID)
	resp := new(types.MoviePaginatedResults)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetReleaseDates retrieves the release date and certification information for a movie.
// Returns release dates and certifications for a movie grouped by country.
// See: https://developer.themoviedb.org/reference/movie-release-dates
func (ms *MoviesService) GetReleaseDates(movieID int) (*types.ReleaseDatesResponse, error) {
	path := fmt.Sprintf("/movie/%d/release_dates", movieID)
	resp := new(types.ReleaseDatesResponse)

	err := ms.Client.DoRequest("GET", path, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// GetReviews retrieves the user reviews for a movie.
// See: https://developer.themoviedb.org/reference/movie-reviews
func (ms *MoviesService) GetReviews(movieID int, opts *LanguagePageOptions) (*types.ReviewPaginatedResults, error) {
	path := fmt.Sprintf("/movie/%d/reviews", movieID)
	resp := new(types.ReviewPaginatedResults)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// GetSimilar retrieves a list of similar movies.
// This is not the same as the "Recommendation" system you see on the website.
// These items are assembled by looking at keywords and genres.
// See: https://developer.themoviedb.org/reference/movie-similar
func (ms *MoviesService) GetSimilar(movieID int, opts *LanguagePageOptions) (*types.MoviePaginatedResults, error) {
	path := fmt.Sprintf("/movie/%d/similar", movieID)
	resp := new(types.MoviePaginatedResults)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTranslations retrieves a list of translations that have been created for a movie.
// See: https://developer.themoviedb.org/reference/movie-translations
func (ms *MoviesService) GetTranslations(movieID int) (*types.TranslationsResponse, error) {
	path := fmt.Sprintf("/movie/%d/translations", movieID)
	resp := new(types.TranslationsResponse)

	err := ms.Client.DoRequest("GET", path, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// GetVideos retrieves the videos that have been added to a movie (trailers, teasers, etc.).
// See: https://developer.themoviedb.org/reference/movie-videos
func (ms *MoviesService) GetVideos(movieID int, opts *LanguageOptions) (*types.VideoList, error) {
	path := fmt.Sprintf("/movie/%d/videos", movieID)
	resp := new(types.VideoList)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// GetWatchProviders retrieves a list of the watch providers (streaming services, rental sources, etc.) for a movie.
// Powered by JustWatch.
// See: https://developer.themoviedb.org/reference/movie-watch-providers
func (ms *MoviesService) GetWatchProviders(movieID int) (*types.WatchProviderResponse, error) {
	path := fmt.Sprintf("/movie/%d/watch/providers", movieID)
	resp := new(types.WatchProviderResponse)

	err := ms.Client.DoRequest("GET", path, nil, nil, resp)
	if err != nil {
		return nil, err
	}

	resp.ID = movieID
	return resp, nil
}

// RateMovie rates a movie.
// A valid session or guest session ID is required.
// See: https://developer.themoviedb.org/reference/movie-add-rating
func (ms *MoviesService) RateMovie(movieID int, body types.RatingRequest, opts *AuthFuncOptions) (*types.StatusResponse, error) {
	if opts != nil && opts.ForGuest {
		if opts.SessionID == nil || opts.GuestSessionID == nil {
			return nil, fmt.Errorf("either SessionID or GuestSessionID is required for RateMovie")
		}
	}

	path := fmt.Sprintf("/movie/%d/rating", movieID)
	resp := new(types.StatusResponse)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ms.Client.DoRequest("POST", path, params, body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeleteRating removes your rating for a movie.
// A valid session or guest session ID is required.
// See: https://developer.themoviedb.org/reference/movie-delete-rating
func (ms *MoviesService) DeleteRating(movieID int, opts *AuthFuncOptions) (*types.StatusResponse, error) {
	if opts != nil && opts.ForGuest {
		if opts.SessionID == nil && opts.GuestSessionID == nil {
			return nil, fmt.Errorf("either SessionID or GuestSessionID is required for DeleteRating")
		}
	}

	path := fmt.Sprintf("/movie/%d/rating", movieID)
	resp := new(types.StatusResponse)
	params, err := utils.StructToURLValues(opts)
	if err != nil {
		return nil, fmt.Errorf("failed to convert options to query params: %w", err)
	}

	err = ms.Client.DoRequest("DELETE", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetLatest retrieves the most recently created movie.
// Note: This endpoint might behave unexpectedly as it returns a single, potentially pre-release movie.
// See: https://developer.themoviedb.org/reference/movie-latest-id
func (ms *MoviesService) GetLatest() (*types.MovieDetails, error) {
	path := "/movie/latest"
	resp := new(types.MovieDetails)

	err := ms.Client.DoRequest("GET", path, nil, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// MovieListOptions represents the available options for list endpoints like NowPlaying, Popular, TopRated, Upcoming.
type MovieListOptions struct {
	Language *string `url:"language,omitempty"`
	Page     *int    `url:"page,omitempty"`
	Region   *string `url:"region,omitempty"` // ISO 3166-1 code
}

// GetNowPlaying retrieves a list of movies that are currently playing in theatres.
// See: https://developer.themoviedb.org/reference/movie-now-playing-list
func (ms *MoviesService) GetNowPlaying(opts *MovieListOptions) (*types.NowPlayingResponse, error) {
	path := "/movie/now_playing"
	resp := new(types.NowPlayingResponse)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetPopular retrieves a list of the current popular movies on TMDb.
// This list updates daily.
// See: https://developer.themoviedb.org/reference/movie-popular-list
func (ms *MoviesService) GetPopular(opts *MovieListOptions) (*types.MoviePaginatedResults, error) {
	path := "/movie/popular"
	resp := new(types.MoviePaginatedResults)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetTopRated retrieves the top rated movies on TMDb.
// See: https://developer.themoviedb.org/reference/movie-top-rated-list
func (ms *MoviesService) GetTopRated(opts *MovieListOptions) (*types.MoviePaginatedResults, error) {
	path := "/movie/top_rated"
	resp := new(types.MoviePaginatedResults)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetUpcoming retrieves a list of movies that are currently scheduled to be released.
// See: https://developer.themoviedb.org/reference/movie-upcoming-list
func (ms *MoviesService) GetUpcoming(opts *MovieListOptions) (*types.UpcomingResponse, error) {
	path := "/movie/upcoming"
	resp := new(types.UpcomingResponse)
	var params url.Values
	var err error
	if opts != nil {
		params, err = utils.StructToURLValues(opts)
		if err != nil {
			return nil, fmt.Errorf("failed to convert options to query params: %w", err)
		}
	}

	err = ms.Client.DoRequest("GET", path, params, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
