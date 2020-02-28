package main

import (
	"localhost.com/go-program-framework/kernel"
	"localhost.com/go-program-framework/utils"
)

// set log
func initLogger() {
	logger = utils.NewLoggerFromMap(logConfig)

	kernel.SetLog(logger)
}
