package config

import (
	"io/ioutil"
	"log"
	"runtime"
	"strings"

	"gopkg.in/yaml.v2"
)

type Conf struct {
	Port       string `yaml:"port"`
	TimeoutMS  int    `yaml:"timeoutMS"`
	MaxAttempt int    `yaml:"maxAttempt"`
}

var config *Conf

func GetConfig() *Conf {
	if config != nil {
		return config
	}
	yamlFile, err := ioutil.ReadFile(getConfigPath())
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalln(err)
	}
	return config
}

func getConfigPath() string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("can't get configuration file path")
	}
	idx := strings.LastIndexByte(file, '/')
	return file[:idx] + "/config.yaml"
}
