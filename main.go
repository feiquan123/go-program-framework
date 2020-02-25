package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	appName    = "your appname"
	appVersion = "v1.0.0"
	v           *viper.Viper
	logger      *logrus.Logger
)

func init() {
	tips()
	initViper()
	printConfig()
	initLogger()

	exit()
}

func main() {
	fmt.Println(v.GetString("log.path"))
	logger.Debug("info")
}
