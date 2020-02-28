package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	appName    = "your appname"
	appVersion = "v1.0.0"
	v          *viper.Viper
	logger     *logrus.Logger
)

func init() {
	tips()
	initViper()
	printConfig()
	initLogger()

	exit()
}

func main() {
	logger.Debug("hello world")
	call()
}
