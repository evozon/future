package future

import (
	"log"
	"strings"

	"github.com/spf13/cobra"

	"future/internal/rector"
)

var AddRuleset = &cobra.Command{
	Use:   "add-ruleset",
	Short: "Adds a rector ruleset",
	Long:  `Edits the rector.php file to add a new ruleset`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isRulesetArgumentValid(args) {
			log.Fatalf("Invalid or missing argument! Example: \\\\Rector\\\\Set\\\\ValueObject\\\\LevelSetList::UP_TO_PHP_81\n")
		}

		file, lines, err := rector.LoadRectorFile()
		if err != nil {
			log.Fatalf(err.Error())
		}

		defer file.Close()

		setsInjectionPoint, err := findLineIndexForSetsMethod(lines)
		if err != nil {
			lines = rector.InjectSetsMethod(lines, args[0])
			if err := rector.WriteRectorFile(file, lines); err != nil {
				log.Fatalf(err.Error())
			}

			return
		}

		lines = rector.InjectLine(lines, setsInjectionPoint, args[0])

		if err := rector.WriteRectorFile(file, lines); err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func isRulesetArgumentValid(args []string) bool {
	// there should be an argument
	if len(args) == 0 {
		return false
	}

	// argument must be a call to a constant
	if !strings.Contains(args[0], "::") {
		return false
	}

	// argument must have a namespace
	if !strings.Contains(args[0], "\\") {
		return false
	}

	return true
}

func findLineIndexForSetsMethod(lines []string) (int, error) {
	index, err := rector.FindLineIndexFor(lines, rector.SetsMethod)

	for err == nil {
		if !strings.Contains(lines[index], "//") {
			return index, err
		}

		index, err = rector.FindLineIndexFor(lines[index:], rector.SetsMethod)
	}

	return index, err
}
