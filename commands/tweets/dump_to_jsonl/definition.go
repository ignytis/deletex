package dump_to_jsonl

import (
	"github.com/integrii/flaggy"
)

func New(inputFile *string, outputFile *string) *flaggy.Subcommand {
	cmd := flaggy.NewSubcommand("tweets:dump:to_jsonl")
	cmd.Description = "Converts a tweet dump file from Javascript to JSON Lines format"
	cmd.String(inputFile, "i", "input-file", "Path to tweets dump file; usually it's data/tweet.js")
	cmd.String(outputFile, "o", "output-file", "Path to JSON Lines output file")

	return cmd
}
