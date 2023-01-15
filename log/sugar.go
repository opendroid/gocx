package log

import (
	"os"

	"go.uber.org/zap"
)

var (
	// Sugar is the global logger for gocx app.
	Sugar *zap.SugaredLogger
)

// init initializes the Sugar logger.
func init() {
	setupSugarLogger()
}

func setupSugarLogger() {
	var config zap.Config
	if os.Getenv("GOCX_ENV") == "DEVELOPMENT" {
		config = zap.NewDevelopmentConfig()
	} else {
		config = zap.NewProductionConfig()
	}
	if logger, err := config.Build(); err != nil {
		panic(err) // OK to panic here
	} else {
		Sugar = logger.Sugar()
	}
}
