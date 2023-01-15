package main

import (
	"os"

	"github.com/opendroid/gocx/cmd"
	"github.com/opendroid/gocx/log"
)

func main() {
	// Execute the root command
	if err := cmd.Execute(); err != nil {
		log.Sugar.Error(err.Error())
		os.Exit(1)
	}
}
