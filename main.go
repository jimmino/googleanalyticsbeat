package main

import (
	"os"

	"github.com/jimmino/googleanalyticsbeat/cmd"

	_ "github.com/jimmino/googleanalyticsbeat/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
