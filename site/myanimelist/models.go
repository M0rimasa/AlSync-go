package myanimelist

const baseURL = "https://myanimelist.net/ownlist"

type animeEntry struct {
	AnimeID            int     `json:"anime_id"`
	NumWatchedEpisodes int     `json:"num_watched_episodes"`
	Status             int     `json:"status"`
	Score              float64 `json:"score"`
	CsrfToken          string  `json:"csrf_token"`
}

type mangaEntry struct {
	MangaID         int     `json:"manga_id"`
	NumReadChapters int     `json:"num_read_chapters"`
	NumReadVolumes  int     `json:"num_read_volumes"`
	Status          int     `json:"status"`
	Score           float64 `json:"score"`
	CsrfToken       string  `json:"csrf_token"`
}