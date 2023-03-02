package cmd

import (
	_ "embed"
	"flag"
	"strings"

	"github.com/opendroid/gocx/dfcx"
	"github.com/opendroid/gocx/log"
)

const (
	// ProjectIDFlag is the projectID flag name.
	_projectIDFlag      = "projectID"
	_projectIDFlagShort = "p"
	// AgentUUIDFlag is the agentID flag name.
	_agentUUIDFlag      = "agentID"
	_agentUUIDFlagShort = "a"
	// VerboseFlag is the verbose flag name.
	_verboseFlag      = "verbose"
	_verboseFlagShort = "v"
)

// getCmdConfig represents the get command config or command line flags.
type getCmdConfig struct {
	projectID string // projectID flag stores the Dialogflow CX project ID in the --projectID flag.
	agentUUID string // agentUUID flag stores the Dialogflow CX agent UUID in the --agentID flag.
	verbose   bool   // verbose flag stores the verbose flag value in the --verbose flag.
}

var (
	// version is the gocx version. value in "gocx --version"
	version = "0.0.1"
	// getFlags is the list of flags for the "root get" command.
	getFlagsSet *flag.FlagSet
	getConfig   getCmdConfig
	// model is the Dialogflow CX model.
	model = dfcx.New()
)

// init initializes the command line flags.
func init() {
	getFlagsSet = flag.NewFlagSet("intents", flag.ExitOnError)
	getFlagsSet.StringVar(&getConfig.projectID, _projectIDFlag, "", "Dialogflow CX project ID")
	getFlagsSet.StringVar(&getConfig.projectID, _projectIDFlagShort, "", "Dialogflow CX project ID (short)")
	getFlagsSet.StringVar(&getConfig.agentUUID, _agentUUIDFlag, "", "Dialogflow CX agent UUID")
	getFlagsSet.StringVar(&getConfig.agentUUID, _agentUUIDFlagShort, "", "Dialogflow CX agent UUID (short)")
	getFlagsSet.BoolVar(&getConfig.verbose, _verboseFlag, false, "verbose output")
	getFlagsSet.BoolVar(&getConfig.verbose, _verboseFlagShort, false, "verbose output (short)")
}

// getIntentsHandler executed on "root get intents" command.
func getIntentsHandler(args []string) {
	// Parse command line flags.
	if err := getFlagsSet.Parse(args); err != nil {
		log.Sugar.Errorw("getIntentsHandler", "error", err.Error())
		return
	}
	if getConfig.verbose {
		log.Sugar.Infow("getIntentsHandler", _verboseFlag, getConfig.verbose, _projectIDFlag, getConfig.projectID, _agentUUIDFlag, getConfig.agentUUID, "n", len(args), "args", strings.Join(args, ", "))
	}
	// Check args.
	if getConfig.agentUUID == "" {
		log.Sugar.Warnw("getIntentsHandler", "error", "missing --agentID flag")
		getFlagsSet.PrintDefaults()
		return
	}
	// projectID, locationID, agentID, languageCode string, verbose bool
	// agentID="9DA50234-36FA-4E09-B10C-C8BE8A424268" random uuid
	model.ShowIntents(getConfig.projectID, "global", getConfig.agentUUID, "en", getConfig.verbose)
}

// getFlowsHandler executed on "root get flows" command.
func getFlowsHandler(args []string) {
	var v bool
	log.Sugar.Infow("getFlowsHandler", "verbose", v, "n", len(args), "args", strings.Join(args, ", "))
}

// getCloseConnection closes the Dialogflow CX clients.
func getCloseConnection() {
	model.Close()
}
