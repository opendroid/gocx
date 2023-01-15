package cmd

import (
	_ "embed"

	"github.com/opendroid/gocx/log"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
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
		Run:     getHandler,
	}

	//go:embed text/getIntentsCmdLong.txt
	getIntentsLong string
	// getIntents is "gocx get intents"
	getIntents = &cobra.Command{
		Use:     "intents",
		Aliases: []string{"i", "intent"},
		Short:   "get the Dialogflow CX agent intents",
		Long:    getIntentsLong,
		Args:    cobra.MinimumNArgs(0),
		Run:     getIntentsHandler,
	}
)

// init adds the getCmd to the rootCmd.
func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(getIntents)
}

func getHandler(cmd *cobra.Command, args []string) {
	for i, arg := range args {
		log.Sugar.Infow("getCmd", zap.Int("index", i), zap.String("arg", arg))
	}
}
func getIntentsHandler(cmd *cobra.Command, args []string) {
	for i, arg := range args {
		log.Sugar.Infow("getIntentsHandler", zap.Int("index", i), zap.String("arg", arg))
	}
}
