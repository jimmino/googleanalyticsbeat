package main

import (
	"os"

	"github.com/elastic/beats/v7/libbeat/logp"
	"github.com/jimmino/googleanalyticsbeat/cmd"

	_ "github.com/jimmino/googleanalyticsbeat/include"
)

func main() {
	logp.Info("Starting GoogleAnalytics Beat... Hit CTRL-C to stop it.")
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
