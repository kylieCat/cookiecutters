package main

import (
	"fmt"
	"flag"

	"github.com/spf13/viper"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/config"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/logger"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/metrics"
	"gitlab.internal.unity3d.com/sre/{{ cookiecutter.svc_name }}/server"
)

var configFileName = flag.String("f", "dev", "Config file name")
var configFilePath = flag.String("d", "/etc", "Config file path")


func readConf(configFilePath, configFileName *string) (*viper.Viper, error) {
	vConf := viper.New()
	vConf.AddConfigPath(*configFilePath)
	vConf.SetConfigType("json")
	vConf.SetConfigName(*configFileName)
	vConf.AutomaticEnv()
	err := vConf.ReadInConfig()
	return vConf, err
}

func main() {
	flag.Parse()

	vConf, err := readConf(configFilePath, configFileName)
	if err != nil {
		panic(fmt.Errorf("could not read in config: %s", err.Error()))
	}
	config.Init(vConf)
	logger.Init()
	metrics.Init()
	server.StartServer()
}
