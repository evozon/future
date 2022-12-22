package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

const (
	defaultPhpVersion = "8.2"
)

var bumpPhp = &cobra.Command{
	Use:   "bump-php",
	Short: "Bump PHP version",
	Long:  `Bump PHP version in composer.json. Must be ran in the root of the project, where the composer.json file is located`,
	Run: func(cmd *cobra.Command, _ []string) {
		phpVersion, err := findPhpVersion()
		if err != nil {
			phpVersion = defaultPhpVersion
		}

		s, file, err := readComposerJson()
		if err != nil {
			log.Fatalf("could not read composer.json: %v\n", err)
		}
		defer file.Close()

		s.setPhpVersion(phpVersion)

		if err := writeComposerJson(file, s); err != nil {
			log.Fatalf("could not write composer.json: %v\n", err)
		}
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
