package myanimelist

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/M0rimasa/AlSync-go/config"
	"github.com/M0rimasa/AlSync-go/site/anilist"
)

type Client struct {
	CsrfToken string
	MALSESSIONID string
	MALHLOGSESSID string
}

func Init(cfg config.Config) Client {
	clientCfg := cfg.MyAnimeList
	return Client {
		clientCfg.CsrfToken,
		clientCfg.Cookies.MALSESSIONID,
		clientCfg.Cookies.MALHLOGSESSID,
	}
}

func (c Client) Update(entry anilist.ListEntry) {
	var url requestURL
	var reqBody []byte

	score := math.Round(entry.Score)
	if entry.Media.Type == "ANIME" {
		malAnimeEntry := animeEntry{
			AnimeID: entry.Media.IDMal,
			NumWatchedEpisodes: entry.Progress,
			Status: convStatus(entry.Status),
			Score: score,
			CsrfToken: c.CsrfToken,
		}
		reqBody, _ = json.Marshal(malAnimeEntry)
		url = requestURL{"anime", "edit"}
	} else {
		malMangaEntry := mangaEntry {
			MangaID: entry.Media.IDMal,
			NumReadChapters: entry.Progress,
			NumReadVolumes: entry.ProgressVolumes,
			Status: convStatus(entry.Status),
			Score: score,
			CsrfToken: c.CsrfToken,
		}
		reqBody, _ = json.Marshal(malMangaEntry)
		url = requestURL{"manga", "edit"}
	}

	cookies := fmt.Sprintf("MALSESSIONID=%v; MALHLOGSESSID=%v; is_logged_in=1;", c.MALSESSIONID, c.MALHLOGSESSID)
	resCode, _ := sendJSON(fmt.Sprint(url), reqBody, cookies)
	if resCode == 500 {
		url.action = "add"
		sendJSON(fmt.Sprint(url), reqBody, cookies)
	}
}


