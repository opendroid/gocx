package cmd

import (
	_ "embed"

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
		Run:     rootCmdHandler,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

// rootCmdHandler executed on root command.
func rootCmdHandler(cmd *cobra.Command, args []string) {
	if v, err := cmd.Flags().GetBool("verbose"); err != nil {
		log.Sugar.Errorw("rootCmdHandler", "error", err)
		return
	} else {
		verbose = v
		log.Sugar.Infow("rootCmdHandler", "verbose", v, "type", v)
	}
}
func init() {
	rootCmd.PersistentFlags().StringVarP(&agentUUID, "agentID", "a", "", "the Dialogflow CX agent UUID")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.MarkFlagRequired("agentID")
}
