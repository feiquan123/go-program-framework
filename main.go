package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	app_name    = "your appname"
	app_version = "v1.0.0"
	v           *viper.Viper
	logger      *logrus.Logger
)

func init() {
	tips()
	initViper()
	initLogger()

	exit()
	//printConfig()
}

func main() {
	fmt.Println(v.GetString("log.path"))
	logger.Debug("info")
}
