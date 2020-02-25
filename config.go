package main

import (
	"log"

	"github.com/spf13/viper"
	"localhost.com/go-program-framework/utils"
)

// config
var (
	logConfig map[string]string
)

// set viper
func initViper() {
	var err error
	v, err = utils.LoadConfig(*c)
	if err != nil {
		log.Fatalln(err)
	}
	readFromConfig(v)
}

func readFromConfig(v *viper.Viper) {
	logConfig = v.GetStringMapString("log")

	if err := utils.Validate(
		logConfig,
	); err != nil {
		log.Fatalln(err)
	}
}

func printConfig() {
	log.Print(
		"\n",
		logConfig, "\n",
		"\n",
	)
}
