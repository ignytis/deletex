package config

type Auth struct {
	ConsumerKey       string `yaml:"consumer_key"`
	ConsumerKeySecret string `yaml:"consumer_key_secret"`
}

type Config struct {
	Auth Auth
}

var config Config
var ConfigPath = "config.yaml"

func GetConfig() *Config {
	return &config
}
