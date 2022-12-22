package main

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var bumpDeps = &cobra.Command{
	Use:   "bump-deps",
	Short: "Bump composer dependencies",
	Long:  `Tries to bump all composer dependencies to the latest version`,
	Run: func(cmd *cobra.Command, _ []string) {
		s, file, err := readComposerJson()
		if err != nil {
			log.Fatalf("could not read composer.json: %v\n", err)
		}
		defer file.Close()

		deps, err := getDepUpgrades()

		updateSchema(deps, &s)

		if err := writeComposerJson(file, s); err != nil {
			log.Fatalf("could not write composer.json: %v\n", err)
		}
	},
}

func updateSchema(deps outdatedDeps, s *schema) {
	// TODO: When trying to bump a dep, do some validation based on the other properties besides Name and Latest
	// TODO: We have a lot of info we can use
	for _, dep := range deps.Installed {
		err := s.setRequireDepVersion(dep.Name, dep.Latest)
		if err != nil {
			err = s.setRequireDevDepVersion(dep.Name, dep.Latest)
			if err != nil {
				log.Printf("could not find dep %s in composer.json. tried both require and require-dev", dep.Name)
				continue
			}
		}
	}

	extra, ok := s.Extra.(map[string]interface{})
	if !ok {
		log.Printf("could not decode extra in composer.json")
		return
	}

	symfony, ok := extra["symfony"].(map[string]interface{})
	if !ok {
		log.Printf("could not decode extra in composer.json")
		return
	}

	_, ok = symfony["require"]
	if !ok {
		log.Printf("could not find require in extra.symfony in composer.json")
		return
	}

	delete(symfony, "require")
	extra["symfony"] = symfony
	s.Extra = extra
}

func getDepUpgrades() (outdatedDeps, error) {
	var deps outdatedDeps

	output, err := exec.Command("composer", "outdated", "--direct", "--format=json").Output()
	if err != nil {
		return deps, err
	}

	err = json.Unmarshal(output, &deps)
	if err != nil {
		return deps, err
	}

	return deps, nil
}

type outdatedDeps struct {
	Installed []struct {
		Name             string  `json:"name"`
		DirectDependency bool    `json:"direct-dependency"`
		Homepage         *string `json:"homepage"`
		Source           string  `json:"source"`
		Version          string  `json:"version"`
		Latest           string  `json:"latest"`
		LatestStatus     string  `json:"latest-status"`
		Description      string  `json:"description"`
		Abandoned        bool    `json:"abandoned"`
		Warning          string  `json:"warning,omitempty"`
	} `json:"installed"`
}
