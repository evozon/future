package future

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"

	"future/internal/composer"
)

const (
	defaultPhpVersion = "8.2"
)

var BumpPhp = &cobra.Command{
	Use:   "bump-php",
	Short: "Bump PHP version",
	Long:  `Bump PHP version in composer.json. Must be run where the composer.json file is located`,
	Run: func(cmd *cobra.Command, _ []string) {
		phpVersion, err := findPhpVersion()
		if err != nil {
			phpVersion = defaultPhpVersion
		}

		s, file, err := composer.ReadComposerJson()
		if err != nil {
			log.Fatalf("could not read composer.json: %v\n", err)
		}

		defer file.Close()

		s.SetPhpVersion(phpVersion)

		if err := composer.WriteComposerJson(file, s); err != nil {
			log.Fatalf("could not write composer.json: %v\n", err)
		}

		log.Printf("successfully bumped the PHP version in the composer.json file to %s\n", phpVersion)
	},
}

func findPhpVersion() (string, error) {
	command := exec.Command("php", "-v")
	out, err := command.Output()
	if err != nil {
		return "", err
	}

	r, err := regexp.Compile(`PHP (\d+\.\d+?\.?\d+)`)
	if err != nil {
		return "", err
	}

	wordAndVersion := r.FindStringSubmatch(string(out))
	if len(wordAndVersion) < 1 {
		return "", fmt.Errorf("failed to find php version in %s\n", string(out))
	}

	return wordAndVersion[1], nil
}
