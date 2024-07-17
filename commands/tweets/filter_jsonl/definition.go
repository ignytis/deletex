package filter_jsonl

import (
	"github.com/integrii/flaggy"
)

func New(inputFile *string, outputFile *string, expression *string) *flaggy.Subcommand {
	cmd := flaggy.NewSubcommand("tweets:jsonl:filter")
	cmd.Description = "Filters a JSON Lines file using expression"
	cmd.String(expression, "e", "expression", "A filtering expression. Check the documentation for examples")
	cmd.String(inputFile, "i", "input-file", "Path to JSON Lines file with tweets. Could be generated using tweets:dump:to_jsonl command")
	cmd.String(outputFile, "o", "output-file", "Path to output file with list of tweet IDs to delete")

	return cmd
}
