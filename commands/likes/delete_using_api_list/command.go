package delete_using_api_list

import (
	"log"

	"github.com/ignytis/deletex/client/twitter"
	"github.com/ignytis/deletex/system/config"
)

func MustRun() {
	log.Println("Deleting likes by requesting a list via API...")

	cfg := config.MustInitialize()
	twitter.MustInitialize(cfg)

	for {
		ls := twitter.ClientInstance.LikesListFirst200()
		for _, l := range ls {
			log.Printf("Deleting like %s...\n", l.IDStr)
			twitter.ClientInstance.LikesMustDelete(l)
		}

		if len(ls) == 0 {
			break
		}
	}

	log.Println("Likes were deleted successfully.")
}
