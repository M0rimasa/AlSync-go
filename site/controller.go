package site

import (
	"fmt"
	"time"

	"github.com/M0rimasa/AlSync-go/config"
	"github.com/M0rimasa/AlSync-go/site/anilist"
	"github.com/M0rimasa/AlSync-go/site/myanimelist"
)

func Updater(c config.Config) {
	fmt.Printf("Keeping an eye out for %v list updates...\n", c.Anilist.Username)
	lastUpdate := time.Now().Unix()
	mal := myanimelist.Init(c)
	for {
		resp, _ := anilist.LastUpdates(c.Anilist.Username)
		entries := resp.Page.MediaList

		for _, entry := range entries {
			if entry.UpdatedAt < lastUpdate {
				continue
			}
				
			fmt.Printf("Updating %v - %v\n", entry.Media.Title.Romaji, entry.Progress)
			mal.Update(entry)
		}
		
		lastUpdate = time.Now().Unix()
		time.Sleep(time.Duration(c.UpdateInterval) * time.Minute)
	}
}
