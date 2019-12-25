package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"localhost.com/ai-go-web-suspicious/utils"
	"log"
)

// tips
var (
	app_name    = "your appname"
	app_version = "v1.0.0"
	v           *viper.Viper
	logger      *logrus.Logger
)

// expend main init
func initConst() {
	initViper()
	initLogger()
}

// set viper
func initViper() {
	var err error
	v, err = utils.LoadConfig(*c)
	if err != nil {
		log.Fatalln(err)
	}
}

// set log
func initLogger() {
	log_level := v.GetString("log.level")
	log_displayLine := v.GetBool("log.displayLine")
	log_isJson := v.GetBool("log.isJson")
	log_console := v.GetBool("log.console")
	log_file := v.GetString("log.file")

	if err := utils.Validate(log_level, log_file); err != nil {
		log.Fatalln(err)
	}
	logger = utils.NewLogger(log_file, log_level, log_isJson, log_console, log_displayLine)
}
