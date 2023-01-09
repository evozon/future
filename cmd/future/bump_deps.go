package future

import (
	"encoding/json"
	"log"
	"os/exec"

	"github.com/spf13/cobra"

	"future/internal/composer"
)

var BumpDeps = &cobra.Command{
	Use:   "bump-deps",
	Short: "Bump composer dependencies",
	Long:  `Bump all Composer dependencies to the latest version. Must be run where the composer.json file is located`,
	Run: func(cmd *cobra.Command, _ []string) {
		s, file, err := composer.ReadComposerJson()
		if err != nil {
			log.Fatalf("could not read composer.json: %v\n", err)
		}

		defer file.Close()

		deps, err := getDepUpgrades()

		updateSchema(deps, &s)

		if err := composer.WriteComposerJson(file, s); err != nil {
			log.Fatalf("could not write composer.json: %v\n", err)
			return
		}

		if len(deps.Installed) == 0 {
			log.Print("all dependencies are at their latest version - nothing to update\n")
			return
		}

		log.Print("successfully updated the following dependencies in the composer.json file:\n")
		for _, dep := range deps.Installed {
			log.Printf("%s: %s -> %s\n", dep.Name, dep.Version, dep.Latest)
		}

		log.Print("run `composer update -W` to apply the changes\n")
	},
}

func updateSchema(deps outdatedDeps, s *composer.Schema) {
	// TODO: When trying to bump a dep, do some validation based on the other properties besides Name and Latest
	// TODO: We have a lot of info we can use
	for _, dep := range deps.Installed {
		err := s.SetRequireDepVersion(dep.Name, dep.Latest)
		if err != nil {
			err = s.SetRequireDevDepVersion(dep.Name, dep.Latest)
			if err != nil {
				log.Printf("could not find dep %s in composer.json. tried both require and require-dev", dep.Name)
				continue
			}
		}
	}

	extra, ok := s.Extra.(map[string]interface{})
	if !ok {
		return
	}

	symfony, ok := extra["symfony"].(map[string]interface{})
	if !ok {
		return
	}

	_, ok = symfony["require"]
	if !ok {
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
