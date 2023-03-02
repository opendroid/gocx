package main

import (
	"os"

	"github.com/opendroid/gocx/cmd"
)

// main is the entry point of the gocx command.
func main() {
	cmd.ExecuteRootCmd(os.Args[:])
}
