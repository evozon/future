package main

import (
	"fmt"
	"os"

	"future/cmd/future"
)

func main() {
	rootCmd := future.RootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
