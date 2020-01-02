package main

import (
	"localhost.com/go-program-framework/utils"
)

// set log
func initLogger() {
	logger = utils.NewLogger(log_file, log_level, log_isJson, log_console, log_displayLine)
}
