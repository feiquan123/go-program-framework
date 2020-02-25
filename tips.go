package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h *bool
	c *string
)

// set default tips
func usage() {
	fmt.Fprintf(os.Stderr, `appname:%s
version:%s
Usage:%s [-h help] [-c configfile]

Options:
`, appName, appVersion, appName)

	flag.PrintDefaults()
}

// parse parameter
func parse() {
	flag.Usage = usage

	if !flag.Parsed() {
		flag.Parse()
	}
	if *h || h == nil {
		flag.Usage()
		os.Exit(0)
	}
	if *c == "" || c == nil {
		fmt.Printf("error: can not get config file form [-c] args.\n\n")
		flag.Usage()
		os.Exit(1)
	}
}

// tips
func tips() {
	h = flag.Bool("h", false, "help")
	c = flag.String("c", "", "config file")

	parse()
}
