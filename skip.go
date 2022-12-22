package main

import (
	"log"
	"strings"

	"github.com/spf13/cobra"
)

var skip = &cobra.Command{
	Use:   "skip",
	Short: "Skips a rector ruleset",
	Long:  `Edits the rector.php file to mark a ruleset as skipped`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isSkipArgumentValid(args) {
			log.Fatalf("Invalid or missing argument! Example: \\\\Rector\\\\Set\\\\ValueObject\\\\LevelSetList::UP_TO_PHP_81::class\n")
		}

		file, lines, err := loadRectorFile()
		if err != nil {
			log.Fatalf(err.Error())
		}

		defer file.Close()

		skipInjectionPoint, err := findLineIndexFor(lines, skipsMethod)
		if err != nil {
			lines = injectSkipMethod(lines, args[0])
			if err := writeRectorFile(file, lines); err != nil {
				log.Fatalf(err.Error())
			}

			return
		}

		lines = injectLine(lines, skipInjectionPoint, args[0])

		if err := writeRectorFile(file, lines); err != nil {
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
