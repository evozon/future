package future

import (
	"log"
	"strings"

	"github.com/spf13/cobra"

	"future/internal/rector"
)

var Skip = &cobra.Command{
	Use:   "skip",
	Short: "Skips a rector ruleset",
	Long:  `Edits the rector.php file to mark a ruleset as skipped`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isSkipArgumentValid(args) {
			log.Fatalf("Invalid or missing argument! Example: \\\\Rector\\\\Set\\\\ValueObject\\\\LevelSetList::UP_TO_PHP_81::class\n")
		}

		file, lines, err := rector.LoadRectorFile()
		if err != nil {
			log.Fatalf(err.Error())
		}

		defer file.Close()

		skipInjectionPoint, err := rector.FindLineIndexFor(lines, rector.SkipsMethod)
		if err != nil {
			lines = rector.InjectSkipMethod(lines, args[0])
			if err := rector.WriteRectorFile(file, lines); err != nil {
				log.Fatalf(err.Error())
			}

			return
		}

		lines = rector.InjectLine(lines, skipInjectionPoint, args[0])

		if err := rector.WriteRectorFile(file, lines); err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func isSkipArgumentValid(args []string) bool {
	// there should be an argument
	if len(args) == 0 {
		return false
	}

	if !strings.Contains(args[0], "::") {
		return false
	}

	// argument must have a namespace
	if !strings.Contains(args[0], "\\") {
		return false
	}

	return true
}
