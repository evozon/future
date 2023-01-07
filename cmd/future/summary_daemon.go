package future

import (
	"net"
	"os"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"

	"future/internal/summary"
)

var summaryDaemon = &cobra.Command{
	Use:   "daemon",
	Short: "Runs a daemon that listens for incoming requests to summarize the output of the upgrade",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Collecting output...")

		runGrpcServer(cmd)
	},
}

func runGrpcServer(cmd *cobra.Command) {
	conn, err := net.Listen("tcp", ":8123")
	if err != nil {
		cmd.PrintErrf("could not open tcp connection: %v", err)
		os.Exit(1)
	}

	var opts []grpc.ServerOption

	srv := grpc.NewServer(opts...)

	summary.RegisterSummaryServer(srv, &summary.Server{})

	if err := srv.Serve(conn); err != nil {
		cmd.PrintErrf("could not start grpc server: %v", err)
		os.Exit(1)
	}
}
