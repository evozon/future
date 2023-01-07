package future

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "future",
		Short: "Future is a binary that helps upgrade projects",
		Long:  `To be used in conjunction with the CI future-proofing stage in projects that require upgrading`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	SummaryCmd.AddCommand(summaryDaemon)
	SummaryCmd.AddCommand(summaryShow)

	rootCmd.AddCommand(BumpPhp)
	rootCmd.AddCommand(BumpDeps)
	rootCmd.AddCommand(AddRule)
	rootCmd.AddCommand(AddRuleset)
	rootCmd.AddCommand(Skip)
	rootCmd.AddCommand(RunCmd)
	rootCmd.AddCommand(SummaryCmd)

	return rootCmd
}
