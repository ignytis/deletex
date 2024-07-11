package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

var ConfigInstance *Config = nil

func MustInitialize() *Config {
	if nil != ConfigInstance {
		return ConfigInstance
	}

	config = Config{}
	yamlFile, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		log.Fatal("Failed to read config file "+ConfigPath+". ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatal("Failed to parse config file "+ConfigPath+". ", err)
	}

	ConfigInstance = &config

	return ConfigInstance
}
