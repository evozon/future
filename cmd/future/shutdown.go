package future

import (
	"fmt"
	"log"
	"sort"

	"github.com/jedib0t/go-pretty/v6/table"
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

		summary := summaryResponse.GetSummary()

		reorder(summary)
		render(summary)

		_, _ = client.Shutdown(cmd.Context(), &collector.ShutdownRequest{})
	},
}

func render(summary []*collector.SummaryData) {
	w := table.NewWriter()
	w.AppendHeader(table.Row{"Command", "Status", "Output"})

	for _, summaryData := range summary {
		result := "successful"
		if summaryData.Status != 0 {
			result = "failed"
		}

		w.AppendRow(table.Row{summaryData.Command, result, summaryData.Output})
	}

	fmt.Println(w.Render())
}

func reorder(summary []*collector.SummaryData) {
	sort.Slice(summary, func(i, j int) bool {
		return summary[i].Status > summary[j].Status
	})
}
