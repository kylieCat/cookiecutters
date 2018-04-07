package main

import (
	"fmt"
	"flag"

	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.svc_name }}/config"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.svc_name }}/logger"
	"github.com/{{ cookiecutter.github_username }}/{{ cookiecutter.svc_name }}/server"
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
		panic(fmt.Errorf("Could not read in config: %s", err.Error()))
	}
	config.Init(vConf)
	logger.Init()
	server.StartServer()
}
