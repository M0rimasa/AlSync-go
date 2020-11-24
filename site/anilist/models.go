package anilist

type ListEntriesResponse struct {
	Page struct {
		MediaList []ListEntry `json:"mediaList"`
	}
}

type ListEntry struct {
	UpdatedAt       int64   `json:"updatedAt"`
	Progress        int     `json:"progress"`
	ProgressVolumes int     `json:"progressVolumes"`
	Score           float64 `json:"score"`
	Status          string  `json:"status"`
	Media           struct {
		ID    int    `json:"id"`
		Type  string `json:"type"`
		IDMal int    `json:"idMal"`
		Title struct {
			Romaji string `json:"romaji"`
		} `json:"title"`
	} `json:"media"`
}