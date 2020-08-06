// Package main is server
//
// main is server entry point
package main

import (
	"flag"
	"log"
)

var configFile = flag.String("c", "configs/app.yaml", "set config file which viper will loading.")
var appName = "your-appname"

func init() {
	flag.Parse()
}

func main() {
	log.Println("app start ....")
	app, err := NewAPP(*configFile)
	if err != nil {
		panic(err)
	}

	if err := app.Start(); err != nil {
		panic(err)
	}

	app.AwaitSingal()
	log.Println("app shutdonw succes")
}
