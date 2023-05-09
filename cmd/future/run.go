package future

import (
	"bytes"
	"fmt"
	"future/internal/collector"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:                "run",
	Short:              "Run any shell command through Future",
	Long:               `Run any shell command through Future`,
	Example:            `bin/future run ls -all`,
	DisableFlagParsing: true,
	Args:               cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var shArgs []string
		var shOut bytes.Buffer

		client, conn := collector.NewClient()
		defer conn.Close()

		if len(args) > 1 {
			shArgs = args[1:]
		}

		sh := exec.Command(args[0], shArgs...)
		sh.Stdout = &shOut

		shCmd := fmt.Sprintf("%s %s", args[0], strings.Join(shArgs, " "))

		if err := sh.Run(); err != nil {
			msg := fmt.Sprintf("failed to run the command: %v\n", err)

			_, err := client.Push(cmd.Context(), &collector.PushRequest{
				Command: shCmd,
				Output:  msg,
				Status:  1,
			})

			if err != nil {
				log.Fatal(msg)
			}

			os.Exit(1)
		}

		_, err := client.Push(cmd.Context(), &collector.PushRequest{
			Command: shCmd,
			Output:  string(shOut.Bytes()),
			Status:  0,
		})

		if err != nil {
			log.Printf("%+v", err)
		}

		os.Exit(0)
	},
}
