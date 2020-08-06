package app_test

import (
	"fmt"

	"github.com/feiquan123/go-program-framework/src/pkg/app"
	"github.com/feiquan123/go-program-framework/src/pkg/config"
	"github.com/feiquan123/go-program-framework/src/pkg/log"
)

// Server implement pkg.app.Server
type Server struct {
	desc string
}

func NewServer() app.Server {
	return &Server{"http、rpc、dns... server"}
}

func (s *Server) Start() error {
	fmt.Println("[START]", s.desc)
	return nil
}

func (s *Server) Stop() error {
	fmt.Println("[STOP]", s.desc)
	return nil
}

func (s *Server) Desc() string {
	return s.desc
}

// ExampleApplicationn create a new application
// you need config[app.yaml]:
func ExampleApplication() {
	/* app.yamls
	app:
		name: your-app-name

	log:
		filename: /tmp/your-app-name.log
		maxSize: 500
		maxBackups: 3
		maxAge: 3
		level: debug
		stdout: true
	*/
	configPath := "./testdata/app.yaml"
	viper, err := config.New(configPath)
	if err != nil {
		panic(err)
	}
	options, err := log.NewOptions(viper)
	if err != nil {
		panic(err)
	}
	logger, err := log.New(options)
	if err != nil {
		panic(err)
	}
	server := NewServer()
	application, err := app.New(viper.GetString("app.name"), logger, server)
	if err != nil {
		panic(err)
	}

	if err := application.Start(); err != nil {
		panic(err)
	}

	application.AwaitSingal()

	// OutPut:
	// 2020/08/06 14:06:28 use local config file -> ./testdata/app.yaml
	// 2020/08/06 14:06:28 read remote config error from 127.0.0.1:7888: Remote Configurations Error: No Files Found
	// 2020/08/06 14:06:28 app config:
	// app:
	// name: your-app-name
	// log:
	// filename: /tmp/your-app-name.log
	// level: debug
	// maxage: 3
	// maxbackups: 3
	// maxsize: 500
	// stdout: true
	//
	// [START] http、rpc、dns... server
	// 2020-08-06T14:06:28.901+0800    INFO    http、rpc、dns... server start success  {"type": "Application"}
	// Ctrl+C
	// 2020-08-06T14:07:15.048+0800  INFO    receive a signal        {"type": "Application", "signal": "interrupt"}
	// [STOP] http、rpc、dns... server
	// 2020-08-06T14:07:15.048+0800    INFO    http、rpc、dns... server shutdown success       {"type": "Application"}
}
