package app

import "go.uber.org/zap"

// TestApplication is run test
type TestApplication struct {
	Application
}

// NewTestApplication is constructor of application
func NewTestApplication(name string, logger *zap.Logger, ss ...Server) (*TestApplication, error) {
	app := Application{
		name:    name,
		logger:  logger.With(zap.String("type", "Application")),
		servers: ss,
	}
	return &TestApplication{app}, nil
}
