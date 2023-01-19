package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/opendroid/gocx/dfcx"
	"github.com/opendroid/gocx/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
		Use:              "gocx",
		Version:          version,
		Short:            "gocx is a command line toolset for managing the Dialogflow CX agents",
		Long:             rootCmdLong,
		PersistentPreRun: rootPFPreRun,
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

func rootPFPreRun(cmd *cobra.Command, args []string) {
	agentUUID = viper.GetString(_agentUUIDFlag)
	verbose = viper.GetBool(_verboseFlag)
	log.Sugar.Infow("rootPFPreRun", _agentUUIDFlag, agentUUID, _verboseFlag, verbose, "n", len(args), "args", args)
}

// init the rootCmd.
func init() {
	cobra.OnInitialize(configViper)
	model = dfcx.New()
	rootCmd.PersistentFlags().StringVarP(&agentUUID, _agentUUIDFlag, "a", "", "the Dialogflow CX agent UUID")
	rootCmd.PersistentFlags().BoolVarP(&verbose, _verboseFlag, "v", false, "verbose output")
	// gotcha: https://stackoverflow.com/questions/52322279/cobra-markpersistentflagrequired-not-working-on-root
	if err := rootCmd.MarkPersistentFlagRequired(_agentUUIDFlag); err != nil {
		log.Sugar.Errorw("init", "error", err)
	}
	viper.BindPFlag(_agentUUIDFlag, rootCmd.PersistentFlags().Lookup(_agentUUIDFlag))
	viper.BindPFlag(_verboseFlag, rootCmd.PersistentFlags().Lookup(_verboseFlag))
	rootCmd.AddCommand(getCmd)
	log.Sugar.Infow("init.configViper", _agentUUIDFlag, viper.GetString(_agentUUIDFlag), _verboseFlag, viper.GetBool(_verboseFlag))
}

// config viper. Enviorment variables are prefixed with "GOCX_" in file .gocx/.env
// and last used command line flags in .gocx/.config.yaml
func configViper() {
	viper.SetEnvPrefix("gocx") // See https://github.com/spf13/cobra/blob/main/user_guide.md
	viper.AutomaticEnv()
	viper.SetConfigName(".config") // config file name without extension, stores last used command line flags
	viper.SetConfigType("yaml")    // Extension
	home, err := os.UserHomeDir()

	cobra.CheckErr(err)
	viper.AddConfigPath(fmt.Sprintf("%s/.gocx", home))
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Sugar.Errorw("configViper", "config", viper.ConfigFileUsed(), "error", err)
		} else {
			log.Sugar.Errorw("configViper", "error", err)
		}
	} else {
		log.Sugar.Infow("configViper", "config", viper.ConfigFileUsed(),
			"home", home, _agentUUIDFlag, viper.Get(_agentUUIDFlag),
			_verboseFlag, viper.Get(_verboseFlag))
	}
}
