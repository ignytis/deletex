package twitter

import (
	"log"

	"github.com/dghubble/go-twitter/twitter"
)

func (c Client) LikesListFirst200() []twitter.Tweet {
	t, _, err := c.Client.Favorites.List(&twitter.FavoriteListParams{
		Count: 200,
	})

	if err != nil {
		log.Fatal("Failed to receive a list of first 200 likes", err)
	}

	return t
}

func (c Client) LikesMustDelete(t twitter.Tweet) {
	_, _, err := c.Client.Favorites.Destroy(&twitter.FavoriteDestroyParams{
		ID: t.ID,
	})
	if err != nil {
		log.Fatal("Failed to delete a like "+t.IDStr, err)
	}
}
