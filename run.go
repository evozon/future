package main

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:                "run",
	Short:              "Run any shell command through the future binary",
	Long:               `Run any shell command through the future binary`,
	Example:            `bin/future run ls -all`,
	DisableFlagParsing: true,
	Args:               cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var shArgs []string

		if len(args) > 1 {
			shArgs = args[1:]
		}

		sh := exec.Command(args[0], shArgs...)

		cmd.Println(sh.String())

		sh.Stdout = cmd.OutOrStdout()
		if err := sh.Run(); err != nil {
			_ = cmd.Help()
			cmd.PrintErrf("Failed to run command: %v\n", err)
			os.Exit(1)
		}
	},
}
