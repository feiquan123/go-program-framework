package main

import (
	"log"

	"github.com/spf13/viper"
	"localhost.com/go-program-framework/kernel"
	"localhost.com/go-program-framework/utils"
)

// config
var (
	logConfig map[string]string
	langurage string
)

func readFromConfig(v *viper.Viper) {
	logConfig = v.GetStringMapString("log")
	langurage = v.GetString("language")

	if err := utils.Validate(
		logConfig, langurage,
	); err != nil {
		log.Fatalln(err)
	}
}

func printConfig() {
	log.Print(
		"\n",
		logConfig, "\n",
		langurage, "\n",
		"\n",
	)
}

// setConfig : set orther model config
func setConfig() {
	kernel.SetConfig(langurage)
}

// initViper : set viper
func initViper() {
	var err error
	v, err = utils.LoadConfig(*c)
	if err != nil {
		log.Fatalln(err)
	}
	readFromConfig(v)
	setConfig()
}
