package future

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"future/internal/collector"
)

var Shutdown = &cobra.Command{
	Use:   "shutdown",
	Short: "Shuts the output collector down",
	Long:  "Shuts down the GRPc server process that accepts output from the commands",
	Run: func(cmd *cobra.Command, _ []string) {
		client, conn := collector.NewClient()

		defer conn.Close()

		summaryResponse, err := client.Summary(cmd.Context(), &collector.SummaryRequest{})
		if err != nil {
			log.Fatalf("could not get the summary: %v\n", err)
		}

		for command, output := range summaryResponse.GetSummary() {
			fmt.Printf("%s -> %s\n", command, output)
		}

		_, _ = client.Shutdown(cmd.Context(), &collector.ShutdownRequest{})
	},
}
