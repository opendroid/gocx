package cmd

import (
	_ "embed"

	"github.com/opendroid/gocx/dfcx"
	"github.com/opendroid/gocx/log"
	"github.com/spf13/cobra"
)

var (
	// version is the gocx version. value in "gocx --version"
	version = "0.0.1"
	// agentUUID flag stores the Dialogflow CX agent UUID in the --agentID flag.
	agentUUID string
	// verbose flag stores the verbose flag value in the --verbose flag.
	verbose bool
	//go:embed text/rootCmdLong.txt
	rootCmdLong string
	rootCmd     = &cobra.Command{
		Use:     "gocx",
		Version: version,
		Short:   "gocx is a command line toolset for managing the Dialogflow CX agents",
		Long:    rootCmdLong,
	}
	// _modelmodel is the interface for the Dialogflow CX model. Used by cmd package.
	model dfcx.Model
)

const (
	// AgentUUIDFlag is the agentID flag name.
	_agentUUIDFlag = "agentID"
	// VerboseFlag is the verbose flag name.
	_verboseFlag = "verbose"
)

func Execute() error {
	return rootCmd.Execute()
}

// init the rootCmd.
func init() {
	model = dfcx.New()
	rootCmd.PersistentFlags().StringVarP(&agentUUID, _agentUUIDFlag, "a", "", "the Dialogflow CX agent UUID")
	rootCmd.PersistentFlags().BoolVarP(&verbose, _verboseFlag, "v", false, "verbose output")
	// gotcha: https://stackoverflow.com/questions/52322279/cobra-markpersistentflagrequired-not-working-on-root
	if err := rootCmd.MarkPersistentFlagRequired(_agentUUIDFlag); err != nil {
		log.Sugar.Errorw("init", "error", err)
	}
}
