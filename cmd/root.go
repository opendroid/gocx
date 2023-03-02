package cmd

import (
	_ "embed"
	"fmt"
)

var (
	//go:embed text/rootCmdLong.txt
	rootCmdText string
	//go:embed text/getCmdLong.txt
	getCmdText string

	// cmdTree Command tree
	cmdTree = map[string][]string{
		"gocx": {"get"},
		"get":  {"intents", "flows"},
	}

	// argsTree Arguments tree
	argsTree = map[string][]string{
		"gocx": {"--version", "-v", "--help", "-h"},
		"get":  {"--agentID", "-a", "--verbose", "-v"},
	}
	_, _ = cmdTree, argsTree
)

// gocx --version
// gocx --help
// gocx get intents --agentID <agentID> --verbose
// gocx get flows --agentID <agentID> --verbose

// ExecuteRootCmd executes the root command.
func ExecuteRootCmd(args []string) {
	switch len(args) {
	case 1: // gocx
		fmt.Println(rootCmdText) // Print the root command help.
	case 2: // gocx <command>
		switch args[1] {
		case "--version":
			fmt.Printf("gocx %s\n", version) // Print the gocx version.
		case "--help", "-h":
			fmt.Println(rootCmdText) // Print the root command help.
		case "get":
			fmt.Println(getCmdText) // Print the get command help..
		case "--agentID", "-a", "--verbose", "-v":
			fmt.Printf("gocx: %s positional argument. Put args in the end\n", args[1])
		default:
			fmt.Printf("gocx: unknown command: %s\n", args[1])
			fmt.Println(rootCmdText) // Print the root command help.
		}
	default: // gocx <command> <subcommand>
		switch args[1] { // Positional argument 1: error
		case "--agentID", "-a", "--verbose", "-v":
			fmt.Printf("gocx: %s positional argument. Put args in the end\n", args[1])
		case "get":
			getIntentsHandler(args[3:])
			getCloseConnection() // Close the Dialogflow CX clients.
		}
	}
}
