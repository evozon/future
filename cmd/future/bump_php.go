package future

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"

	"future/internal/collector"
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
		client, conn := collector.NewClient()
		defer conn.Close()

		phpVersion, err := findPhpVersion()
		if err != nil {
			phpVersion = defaultPhpVersion
		}

		s, file, err := composer.ReadComposerJson()
		if err != nil {
			msg := fmt.Sprintf("could not read composer.json: %v\n", err)
			_, err := client.Push(cmd.Context(), &collector.PushRequest{
				Command: cmd.Name(),
				Output:  msg,
				Status:  1,
			})

			if err != nil {
				log.Fatal(msg)
			}

			os.Exit(1)
		}

		defer file.Close()

		s.SetPhpVersion(phpVersion)

		if err := composer.WriteComposerJson(file, s); err != nil {
			msg := fmt.Sprintf("could not write composer.json: %v\n", err)

			_, err := client.Push(cmd.Context(), &collector.PushRequest{
				Command: cmd.Name(),
				Output:  msg,
				Status:  1,
			})

			if err != nil {
				log.Fatal(msg)
			}

			os.Exit(1)
		}

		msg := fmt.Sprintf("successfully bumped PHP version in composer.json to %s\n", phpVersion)
		_, err = client.Push(cmd.Context(), &collector.PushRequest{
			Command: cmd.Name(),
			Output:  msg,
			Status:  0,
		})

		if err != nil {
			log.Printf("%+v", err)
		}

		os.Exit(0)
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
