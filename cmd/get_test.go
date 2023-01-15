package cmd

import (
	_ "embed"
	"fmt"
	"testing"
)

var (
	//go:embed text/getCmd_test.txt
	getCmdTest string
)

// TestExecute_GetCmd tests the root command.
// go test -run TestExecute_GetCmd -v ./...
func TestExecute_GetCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  error
		want string
	}{
		{name: "get: missing subcommand", args: []string{"get"}, err: nil, want: getCmdTest},
		{name: "get intents: missing parameter", args: []string{"get", "intents"}, err: fmt.Errorf(`required flag(s) "agentID" not set`)},
		{name: "get intents: all params", args: []string{"get", "intents", `--agentID="test_uuid"`}, err: nil},
		{name: "get flows: missing parameter", args: []string{"get", "flows"}, err: fmt.Errorf(`required flag(s) "agentID" not set`)},
		{name: "get flows: all params", args: []string{"get", "flows", `--agentID="test_uuid"`}, err: nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := testExecuteCommand(t, rootCmd, test.args)
			if err != nil {
				if test.err != nil && err.Error() != test.err.Error() { // Unexpected error
					t.Errorf("rootCmd.Execute() error = %s, wantErr %s", err.Error(), test.err.Error())
				}
				return
			}
			if got != test.want {
				t.Errorf("rootCmd.Execute():\ngot:===>\n%swant:===>\n%s", got, test.want)
			}
		})
	}
}
