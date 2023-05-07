package future

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"future/internal/collector"
)

var Collect = &cobra.Command{
	Use:   "collect",
	Short: "Setup the output collector",
	Long:  "Starts a GRPc server process that accepts output from the various commands in the CI process",
	Run: func(cmd *cobra.Command, _ []string) {
		srv, err := collector.NewServer()
		if err != nil {
			log.Fatalf("could not create collector server: %v\n", err)
		}

		fmt.Printf("listening on %s\n", collector.Address)
		if err := srv.Start(); err != nil {
			srv.Stop()
			log.Fatalf("could not start collector server: %v\n", err)
		}
	},
}
