package types

// GroupedEpisode represents an episode within an episode group.
// It extends TVEpisodeListResult with an 'order' field.
type GroupedEpisode struct {
	TVEpisodeListResult         // Embed basic episode info
	Order               int `json:"order"` // Order of the episode within this group
}

// EpisodeGroupDetails represents the detailed information for a TV episode group.
// See: https://developer.themoviedb.org/reference/tv-episode-group-details
type EpisodeGroupDetails struct {
	Description  string           `json:"description"`
	EpisodeCount int              `json:"episode_count"`
	GroupCount   int              `json:"group_count"`
	Groups       []GroupedEpisode `json:"groups"` // List of episodes in this specific group
	ID           string           `json:"id"` // Group ID (string)
	Name         string           `json:"name"`
	Network      *Network         `json:"network,omitempty"` // Nullable, uses Network from networks.go
	Type         int              `json:"type"` // Type identifier
}

