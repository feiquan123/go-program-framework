package main

import (
	"os"
	"os/signal"
	"syscall"
)

func exit() {

	go func() {
		var (
			c chan os.Signal
			s os.Signal
		)

		c = make(chan os.Signal, 1)
		signal.Reset(syscall.SIGTERM, syscall.SIGINT)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

		for {
			s = <-c
			logger.Info("+++++++++++++++ 收到信号,退出程序", s.Signal, " +++++++++++++++++++")
			// add exit program

			logger.Info("+++++++++++++++ app exit success +++++++++++++++++++++++")
			os.Exit(0)
		}
	}()
}
