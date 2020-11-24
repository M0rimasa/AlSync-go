package anilist

import (
	"context"

	"github.com/machinebox/graphql"
)

var client = graphql.NewClient("https://graphql.anilist.co")

func LastUpdates(username string) (response ListEntriesResponse, err error) {
	req := graphql.NewRequest(`
    query ($perPage: Int, $user: String) {
	  Page(perPage: $perPage) {
	    mediaList(userName: $user, sort: UPDATED_TIME_DESC, status_not: REPEATING) {
	      updatedAt
	      progress
	      progressVolumes
	      score(format: POINT_10_DECIMAL)
	      status
	      media {
	        id
	        type
	        idMal
	        title {
	          romaji
	        }
	      }
	    }
	  }
	}
	`)
	req.Var("user", username)
	req.Var("perPage", 30)
	err = client.Run(context.Background(), req, &response)
	return
}