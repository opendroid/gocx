package cmd

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/opendroid/gocx/log"
	"github.com/spf13/cobra"
)

var (
	//go:embed text/getCmdLong.txt
	getCmdLong string
	// getCmd is "gocx get"
	getCmd = &cobra.Command{
		Use:     "get",
		Aliases: []string{"g"},
		Short:   "get the Dialogflow CX agent intents or flows or pages",
		Long:    getCmdLong,
		Args:    cobra.MinimumNArgs(1),
	}

	//go:embed text/getIntentsCmdLong.txt
	getIntentsLong string
	// getIntents is "gocx get intents"
	getIntents = &cobra.Command{
		Use:     "intents",
		Aliases: []string{"i", "intent"},
		Short:   "get the Dialogflow CX agent intents",
		Long:    getIntentsLong,
		Args:    cobra.ExactArgs(0),
		Run:     getIntentsHandler,
	}

	//go:embed text/getFlowsCmdLong.txt
	getFlowsLong string
	// getFlows is "gocx get flows"
	getFlows = &cobra.Command{
		Use:     "flows",
		Aliases: []string{"f", "flow"},
		Short:   "get the Dialogflow CX agent flows",
		Long:    getFlowsLong,
		Args:    cobra.ExactArgs(0),
		Run:     getFlowsHandler,
	}
)

// init the getCmd and it as sub-command of rootCmd: root => get => intents
func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getIntents, getFlows)
}

// getIntentsHandler executed on "root get intents" command.
func getIntentsHandler(cmd *cobra.Command, args []string) {
	var v bool
	var err error
	if v, err = cmd.Flags().GetBool("verbose"); err != nil {
		log.Sugar.Errorw("getIntentsHandler", "error", err, "n", len(args), "args", strings.Join(args, ", "))
		return
	}
	log.Sugar.Infow("getIntentsHandler", "verbose", v, "type", fmt.Sprintf("%T", v), "n", len(args), "args", strings.Join(args, ", "))
}

// getFlowsHandler executed on "root get flows" command.
func getFlowsHandler(cmd *cobra.Command, args []string) {
	var v bool
	var err error
	if v, err = cmd.Flags().GetBool("verbose"); err != nil {
		log.Sugar.Errorw("getFlowsHandler", "error", err, "n", len(args), "args", strings.Join(args, ", "))
		return
	}
	log.Sugar.Infow("getFlowsHandler", "verbose", v, "type", fmt.Sprintf("%T", v), "n", len(args), "args", strings.Join(args, ", "))
}
