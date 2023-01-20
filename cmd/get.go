package cmd

import (
	_ "embed"
	"strings"

	"github.com/opendroid/gocx/dfcx"
	"github.com/opendroid/gocx/log"
)

const (
	// AgentUUIDFlag is the agentID flag name.
	_agentUUIDFlag = "agentID"
	// VerboseFlag is the verbose flag name.
	_verboseFlag = "verbose"
)

var (
	// version is the gocx version. value in "gocx --version"
	version = "0.0.1"
	// agentUUID flag stores the Dialogflow CX agent UUID in the --agentID flag.
	agentUUID string
	// verbose flag stores the verbose flag value in the --verbose flag.
	verbose bool
	_, _    = agentUUID, verbose // to avoid "declared and not used" error.
	model   = dfcx.New()
)

// getIntentsHandler executed on "root get intents" command.
func getIntentsHandler(args []string) {
	var v bool

	log.Sugar.Infow("getIntentsHandler", _verboseFlag, v, _agentUUIDFlag, agentUUID, "n", len(args), "args", strings.Join(args, ", "))
	// projectID, locationID, agentID, languageCode string, verbose bool
	// agentID="20481884-1290-479c-9e66-8b8c72e3f6eb"
	model.ShowIntents("eats-universal-a", "global", agentUUID, "en", v)
}

// getFlowsHandler executed on "root get flows" command.
func getFlowsHandler(args []string) {
	var v bool

	log.Sugar.Infow("getFlowsHandler", "verbose", v, "n", len(args), "args", strings.Join(args, ", "))
}
