package main

import (
	"github.com/spf13/viper"
	"localhost.com/go-program-framework/utils"
	"log"
)

// config
var (
	log_level       string
	log_displayLine bool
	log_isJson      bool
	log_console     bool
	log_file        string
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
	log_level = v.GetString("log.level")
	log_displayLine = v.GetBool("log.displayLine")
	log_isJson = v.GetBool("log.isJson")
	log_console = v.GetBool("log.console")
	log_file = v.GetString("log.file")

	if err := utils.Validate(
		log_level, log_file,
	); err != nil {
		log.Fatalln(err)
	}
}

func printConfig() {
	log.Println(
		"\n",
		log_level, "\n",
		log_displayLine, "\n",
		log_isJson, "\n",
		log_console, "\n",
		log_file, "\n",
	)
}
