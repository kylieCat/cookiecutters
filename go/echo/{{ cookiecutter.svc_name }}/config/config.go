package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

var vConf *viper.Viper
var tz *time.Location

func Init(conf *viper.Viper) error {
	if conf == nil {
		return fmt.Errorf("nil config passed")
	} else {
		setConfig(conf)
	}

	if t := GetConfig().GetString("timezone"); t != "" {
		var err error
		tz, err = time.LoadLocation(t)
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}

func GetTimezone() *time.Location {
	return tz
}

func setConfig(conf *viper.Viper) {
	vConf = conf
}

func GetConfig() *viper.Viper {
	return vConf
}
