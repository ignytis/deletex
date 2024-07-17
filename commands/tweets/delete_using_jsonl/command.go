package delete_using_jsonl

import (
	"bufio"
	"encoding/json"
	"log"
	"os"

	"github.com/ignytis/deletex/client/twitter"
	"github.com/ignytis/deletex/system/config"
	"github.com/ignytis/deletex/types"
)

func MustRun(inputFile string) {
	log.Println("Deleting tweets using file '" + inputFile + "'...")

	in, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("Unable to read input file "+inputFile+". ", err)
	}
	defer in.Close()
	scanner := bufio.NewScanner(in)

	cfg := config.MustInitialize()
	twitter.MustInitialize(cfg)

	for scanner.Scan() {
		line := scanner.Text()
		var record types.Record
		err = json.Unmarshal([]byte(line), &record)
		if err != nil {
			log.Fatal("Unable to read marshal a line: "+line+". ", err)
		}

		log.Printf("Deleting tweet %s...\n", record.Tweet.Id)
		twitter.ClientInstance.MustDeleteTweet(record.Tweet.Id)
	}

	log.Println("Tweets were deleted successfully.")
}
