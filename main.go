package main

import (
	"log"

	"github.com/integrii/flaggy"

	likes_delete_using_api_list "github.com/ignytis/deletex/commands/likes/delete_using_api_list"
	"github.com/ignytis/deletex/commands/tweets/delete_using_jsonl"
	"github.com/ignytis/deletex/commands/tweets/dump_to_jsonl"
	"github.com/ignytis/deletex/commands/tweets/filter_jsonl"
	"github.com/ignytis/deletex/system/config"
)

func init() {
	flaggy.DefaultParser.ShowHelpOnUnexpected = true
	flaggy.DefaultParser.ShowHelpWithHFlag = true
}

func main() {
	var inputFile, outputFile, expression = "", "", ""

	flaggy.SetName("deletex")
	flaggy.SetDescription(`A utility for tweet deletion. Webpage: https://github.com/ignytis/deletex`)

	flaggy.String(&config.ConfigPath, "c", "config", "Configuration file. See config.example.yaml for more details.")

	subcommandLikesDeleteUsingApiList := likes_delete_using_api_list.New()
	flaggy.AttachSubcommand(subcommandLikesDeleteUsingApiList, 1)

	subcommandTweetsDumpToJsonl := dump_to_jsonl.New(&inputFile, &outputFile)
	flaggy.AttachSubcommand(subcommandTweetsDumpToJsonl, 1)

	subcommandTweetsDeleteUsingJsonl := delete_using_jsonl.New(&inputFile)
	flaggy.AttachSubcommand(subcommandTweetsDeleteUsingJsonl, 1)

	subcommandTweetsToDeleteListFromJsonl := filter_jsonl.New(&inputFile, &outputFile, &expression)
	flaggy.AttachSubcommand(subcommandTweetsToDeleteListFromJsonl, 1)

	flaggy.Parse()

	log.Println("Starting...")
	if subcommandLikesDeleteUsingApiList.Used {
		likes_delete_using_api_list.MustRun()
	} else if subcommandTweetsDumpToJsonl.Used {
		dump_to_jsonl.MustRun(inputFile, outputFile)
	} else if subcommandTweetsDeleteUsingJsonl.Used {
		delete_using_jsonl.MustRun(inputFile)
	} else if subcommandTweetsToDeleteListFromJsonl.Used {
		filter_jsonl.MustRun(inputFile, outputFile, expression)
	}

	log.Println("Done.")
}
