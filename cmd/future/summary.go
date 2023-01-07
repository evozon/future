package future

import (
	"github.com/spf13/cobra"
)

var SummaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Exposes 2 commands for summarizing the output of the upgrade",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}
