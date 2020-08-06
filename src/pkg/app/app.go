package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Server define register server interface
type Server interface {
	Desc() string
	Start() error
	Stop() error
}

// Application define application , run on test,aliyun,online
type Application struct {
	name    string
	logger  *zap.Logger
	servers []Server
}

// New is constructor of application
func New(name string, logger *zap.Logger, ss ...Server) (*Application, error) {
	app := &Application{
		name:    name,
		logger:  logger.With(zap.String("type", "Application")),
		servers: ss,
	}
	return app, nil
}

// Start application
func (a *Application) Start() error {
	if len(a.servers) != 0 {
		for _, s := range a.servers {
			if s != nil {
				desc := s.Desc()
				if err := s.Start(); err != nil {
					return errors.Wrap(err, fmt.Sprintf("%s start errors", desc))
				}
				a.logger.Info(fmt.Sprintf("%s start success", desc))
			}
		}
	}

	return nil
}

// AwaitSingal graceful shutodwn when receive termainal signal
// kill -15
func (a *Application) AwaitSingal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		a.logger.Info("receive a signal", zap.String("signal", s.String()))
		if len(a.servers) != 0 {
			for _, s := range a.servers {
				if s != nil {
					desc := s.Desc()
					if err := s.Stop(); err != nil {
						a.logger.Error(fmt.Sprintf("%s shutdown errors,err:%v", desc, err))
						return
					}
					a.logger.Info(fmt.Sprintf("%s shutdown success", desc))
				}
			}
		}
		os.Exit(0)
	}
}
