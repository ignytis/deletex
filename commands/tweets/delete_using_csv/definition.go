package delete_using_csv

import (
	"github.com/integrii/flaggy"
)

func New(inputFile *string) *flaggy.Subcommand {
	cmd := flaggy.NewSubcommand("tweets:delete:using_csv")
	cmd.Description = "Deletes tweets using a CSV file as a data source"
	cmd.String(inputFile, "i", "input-file", "Path to CSV file where each line is ID of tweet to delete")

	return cmd
}
