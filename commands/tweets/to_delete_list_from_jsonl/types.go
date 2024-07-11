package to_delete_list_from_jsonl

import (
	"log"
	"strconv"
	"time"
)

type Tweet struct {
	CreatedAtString     string `json:"created_at"`
	Id                  string `json:"id"`
	FavoriteCountString string `json:"favorite_count"`
	FullText            string `json:"full_text"`
	RetweetCountString  string `json:"retweet_count"`
}

func (t Tweet) CreatedAt() time.Time {
	timeT, err := time.Parse("Mon Jan 2 15:04:05 MST 2006", t.CreatedAtString)
	if err != nil {
		log.Fatal("Failed to convert the Creted At param '"+t.CreatedAtString+"' to time: ", err)
	}

	timeT.Format("2006-01-02 15:04:05")

	return timeT
}

func (t Tweet) FavoriteCount() int {
	intVar, err := strconv.Atoi(t.FavoriteCountString)
	if err != nil {
		log.Fatal("Failed to convert the Favorite Count param '"+t.FavoriteCountString+"' to integer: ", err)
	}

	return intVar
}

func (t Tweet) RetweetCount() int {
	intVar, err := strconv.Atoi(t.RetweetCountString)
	if err != nil {
		log.Fatal("Failed to convert the Retweet Count param '"+t.RetweetCountString+"' to integer: ", err)
	}

	return intVar
}

type Record struct {
	Tweet Tweet `json:"tweet"`
}
