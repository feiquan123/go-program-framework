package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

// load config
func LoadConfig(file string) (v *viper.Viper, err error) {
	file, err = filepath.Abs(file)
	if err != nil {
		return nil, err
	}
	if _, err := os.Open(file); err != nil {
		panic(fmt.Sprintf("can not read config file, %s", file))

	}

	filelist := strings.Split(filepath.Base(file), ".")
	filename := strings.Join(filelist[0:len(filelist)-1], ".")
	filesuffix := filelist[len(filelist)-1]

	// viper
	v = viper.New()
	v.AddConfigPath(filepath.Dir(file))
	v.SetConfigName(filename)
	v.SetConfigType(filesuffix)

	if err = v.ReadInConfig(); err != nil {
		return nil, err
	}
	return v, nil
}
