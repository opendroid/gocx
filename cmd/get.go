package cmd

import (
	_ "embed"
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
	if v, err = cmd.Flags().GetBool(_verboseFlag); err != nil {
		log.Sugar.Errorw("getIntentsHandler", "error", err, "flag", _verboseFlag, "n", len(args), "args", strings.Join(args, ", "))
		return
	}
	if agentUUID, err = cmd.Flags().GetString(_agentUUIDFlag); err != nil {
		log.Sugar.Errorw("getIntentsHandler", "error", err, "flag", _agentUUIDFlag, "n", len(args), "args", strings.Join(args, ", "))
		return
	}
	log.Sugar.Infow("getIntentsHandler", _verboseFlag, v, _agentUUIDFlag, agentUUID, "n", len(args), "args", strings.Join(args, ", "))
	// projectID, locationID, agentID, languageCode string, verbose bool
	// agentID="20481884-1290-479c-9e66-8b8c72e3f6eb"
	model.ShowIntents("eats-universal-a", "global", agentUUID, "en", v)
}

// getFlowsHandler executed on "root get flows" command.
func getFlowsHandler(cmd *cobra.Command, args []string) {
	var v bool
	var err error
	if v, err = cmd.Flags().GetBool(_verboseFlag); err != nil {
		log.Sugar.Errorw("getFlowsHandler", "error", err, "n", len(args), "args", strings.Join(args, ", "))
		return
	}
	log.Sugar.Infow("getFlowsHandler", "verbose", v, "n", len(args), "args", strings.Join(args, ", "))
}
