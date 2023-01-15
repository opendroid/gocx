package cmd

import (
	"bytes"
	_ "embed"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
)

// TestExecute_RootCmd tests the root command.
// See https://nayaktapan37.medium.com/testing-cobra-commands-in-golang-ca1fe4ad6657
// go test -run TestExecute_RootCmd -v ./...
func TestExecute_RootCmd(t *testing.T) {
	tests := []struct {
		name string
		args []string
		err  error
		want string
	}{
		{name: "invalid command", args: []string{"nada"}, err: fmt.Errorf(`unknown command "nada" for "gocx"`)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := testExecuteCommand(t, rootCmd, test.args)
			if err != nil {
				if test.err != nil && err.Error() != test.err.Error() { // Unexpected error
					t.Errorf("TestExecute_RootCmd: gotError: %s, wantErr: %s", err.Error(), test.err.Error())
				}
				return
			}
			if got != test.want {
				t.Errorf("TestExecute_RootCmd:\ngot:===>\n%swant:===>\n%s", got, test.want)
			}
		})
	}
}

// testExecuteCommand executes a command and returns the output and error.
func testExecuteCommand(t *testing.T, cmd *cobra.Command, args []string) (string, error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	err := cmd.Execute()
	return buf.String(), err
}
