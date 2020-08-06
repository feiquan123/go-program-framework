package config

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/google/wire"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"gopkg.in/yaml.v2"
)

// New define constructor *viper.Viper
func New(path string) (v *viper.Viper, err error) {
	v, err = readConfig(path)
	if err != nil {
		return nil, err
	}

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("config[%s] changed :%s", path, e.Name)
		vt, err := readConfig(path)
		if err != nil {
			log.Printf("[E:] readConfig error, %s", err.Error())
		}
		v = vt
	})

	return v, nil
}

// readConfig is read config
func readConfig(path string) (v *viper.Viper, err error) {
	v = viper.New()

	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err == nil {
		log.Printf("use local config file -> %s\n", v.ConfigFileUsed())
	} else {
		return nil, err
	}

	ys, err := yaml2String(v)
	if err != nil {
		return nil, err
	}
	log.Printf("app config:\n%v\n", ys)
	return v, err
}

func yaml2String(v *viper.Viper) (string, error) {
	c := v.AllSettings()
	bs, err := yaml.Marshal(c)
	if err != nil {
		return "", errors.Wrap(err, "config marshal to string error")
	}
	return string(bs), nil
}

// ProviderSet is wire provider set of config
var ProviderSet = wire.NewSet(New)
