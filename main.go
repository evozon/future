package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd.AddCommand(bumpPhp)
	rootCmd.AddCommand(bumpDeps)
	rootCmd.AddCommand(addRule)
	rootCmd.AddCommand(addRuleset)
	rootCmd.AddCommand(skip)

	Execute()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "future",
	Short: "Future is a binary that helps upgrade projects",
	Long:  `To be used in conjunction with the gitlab future-proofing stage in projects that require upgrading`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: describe the commands here
	},
}
