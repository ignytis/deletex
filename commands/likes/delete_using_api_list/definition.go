package delete_using_api_list

import (
	"github.com/integrii/flaggy"
)

func New() *flaggy.Subcommand {
	cmd := flaggy.NewSubcommand("likes:delete:using_api_list")
	cmd.Description = "Deletes likes using API to query a list"

	return cmd
}
