package types

// FindResponse represents the results from the find by external ID endpoint.
// It can contain results for movies, TV shows, people, TV seasons, and TV episodes.
// See: https://developer.themoviedb.org/reference/find-by-id
type FindResponse struct {
	MovieResults     []MovieListResult     `json:"movie_results"`
	TVResults        []TVListResult        `json:"tv_results"`
	PersonResults    []PersonListResult    `json:"person_results"`
	TVEpisodeResults []TVEpisodeListResult `json:"tv_episode_results"`
	TVSeasonResults  []TVSeasonListResult  `json:"tv_season_results"`
}

