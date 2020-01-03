package main

import (
	"localhost.com/go-program-framework/utils"
)

// set log
func initLogger() {
	logger = utils.NewLoggerFromMap(log_config)
}
