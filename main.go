package main

import (
	"fmt"
)

func init() {
	tips()
	initConst()
}

func main() {
	fmt.Println(v.GetString("log.path"))
	logger.Debug("info")
}
