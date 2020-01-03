package main

import (
	"github.com/spf13/viper"
	"localhost.com/go-program-framework/utils"
	"log"
)

// config
var (
	log_config map[string]string
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
	log_config = v.GetStringMapString("log")

	if err := utils.Validate(
		log_config,
	); err != nil {
		log.Fatalln(err)
	}
}

func printConfig() {
	log.Println(
		"\n",
		log_config, "\n",
	)
}
