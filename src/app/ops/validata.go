package ops

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/feiquan123/go-program-framework/src/pkg/utils"
)

func validate(v *viper.Viper) error {
	if err := utils.Validate(v.GetString("app.name")); err != nil {
		return errors.Wrap(err, "config app.name error")
	}

	return nil
}
