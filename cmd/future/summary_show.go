package future

import (
	"os"

	"github.com/spf13/cobra"

	"future/internal/summary"
)

var summaryShow = &cobra.Command{
	Use:   "show",
	Short: "Shows the summary of the output of the upgrade",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Showing summary...")

		c, conn, err := summary.NewClient()
		if err != nil {
			cmd.PrintErrf("could not create client: %v", err)
			os.Exit(1)
		}
		defer conn.Close()

		stream, err := c.Pull(cmd.Context(), &summary.PullRequest{})
		if err != nil {
			cmd.PrintErrf("could not pull: %v", err)
			os.Exit(1)
		}

		cmd.Println(stream.Recv())
	},
}
