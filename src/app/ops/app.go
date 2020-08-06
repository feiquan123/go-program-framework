package ops

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/feiquan123/go-program-framework/src/pkg/app"
)

// NewAPP creates a new instance of the Application
func NewAPP(v *viper.Viper, logger *zap.Logger) (a *app.Application, err error) {
	if err := validate(v); err != nil {
		return nil, err
	}

	a, err = app.New(v.GetString("app.name"), logger)
	if err != nil {
		return nil, err
	}

	return a, nil
}
