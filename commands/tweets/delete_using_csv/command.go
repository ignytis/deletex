package delete_using_csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/ignytis/deletex/client/twitter"
	"github.com/ignytis/deletex/system/config"
)

func MustRun(inputFile string) {
	log.Println("Deleting tweets using file '" + inputFile + "'...")

	f, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Unable to read input file "+inputFile+". ", err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)

	cfg := config.MustInitialize()
	twitter.MustInitialize(cfg)

	for {
		var tweetIdLine, err = csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Failed to read a line from file "+inputFile+". ", err)
		}

		log.Printf("Deleting tweet %s...\n", tweetIdLine[0])
		twitter.ClientInstance.MustDeleteTweet(tweetIdLine[0])
	}

	log.Println("Tweets were deleted successfully.")
}
