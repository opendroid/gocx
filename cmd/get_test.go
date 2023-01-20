package cmd

import (
	_ "embed"
	"testing"
)

var (
	//go:embed text/getCmd_test.txt
	getCmdTest string
)

// TestExecute_GetCmd tests the root command.
// go test -run TestExecute_GetCmd -v ./...
func TestExecute_GetCmd(t *testing.T) {
	getIntentsHandler([]string{"get", "intents"})
}
