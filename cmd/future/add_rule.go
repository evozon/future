package future

import (
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	"future/internal/rector"
)

var AddRule = &cobra.Command{
	Use:   "add-rule",
	Short: "Adds a rector rule",
	Long:  `Edits the rector.php file to add a new rule`,
	Run: func(cmd *cobra.Command, args []string) {
		if !isRuleArgumentValid(args) {
			log.Fatalf("Invalid or missing argument! Example: \\\\Rector\\\\Set\\\\ValueObject\\\\LevelSetList::UP_TO_PHP_81::class\n")
		}

		file, lines, err := rector.LoadRectorFile()
		if err != nil {
			log.Fatalf(err.Error())
		}
		defer file.Close()

		ruleInjectionPoint, err := rector.FindLineIndexFor(lines, rector.RulesMethod)
		if err != nil {
			// if we can't find a ->rules([...]) section, we'll try to find a ->rule(...) section and convert it to a ->rules section
			lines = convertSingleRuleToMultipleRules(lines)

			ruleInjectionPoint, err = rector.FindLineIndexFor(lines, rector.RulesMethod)
			if err != nil {
				log.Fatalf(err.Error())
			}
		}

		lines = rector.InjectLine(lines, ruleInjectionPoint, args[0])

		if err := rector.WriteRectorFile(file, lines); err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func convertSingleRuleToMultipleRules(lines []string) []string {
	for index, line := range lines {
		if !strings.Contains(line, rector.RuleMethod) {
			continue
		}

		indentation := line[:strings.Index(line, rector.RuleMethod)]

		lines[index] = fmt.Sprintf("%s%s([", indentation, rector.RulesMethod)
		lines = append(
			lines[:index+1],
			append(
				[]string{fmt.Sprintf("%s]);", indentation)},
				lines[index+1:]...,
			)...,
		)

		if strings.Contains(lines[index-1], "register a single rule") {
			lines[index-1] = fmt.Sprintf("%s// register multiple rules", indentation)
		}
	}

	return lines
}

func isRuleArgumentValid(args []string) bool {
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
