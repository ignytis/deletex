package delete_using_jsonl

import (
	"github.com/integrii/flaggy"
)

func New(inputFile *string) *flaggy.Subcommand {
	cmd := flaggy.NewSubcommand("tweets:delete:using_jsonl")
	cmd.Description = "Deletes tweets using JSON Lines file as a data source"
	cmd.String(inputFile, "i", "input-file", "Path to JSON Lines file")

	return cmd
}
