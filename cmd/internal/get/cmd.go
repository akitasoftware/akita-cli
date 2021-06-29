package get

import (
	"github.com/spf13/cobra"
)

// Parent command for listing objects from Akita (just traces for now.)
var Cmd = &cobra.Command{
	Use:          "list",
	Short:        "List objects in the Akita cloud.",
	Long:         "List objects in the Akita cloud.",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}