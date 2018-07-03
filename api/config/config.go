package config

import (
	"github.com/jinzhu/configor"
)

// Config : config info
type Config struct {
	AppName string `default:"zmemo"`
	Port    string `default:"8080"`
	DBLog   bool   `default:false`
	DB      struct {
		Name     string `default:"zmemo"`
		User     string `default:"zmemo"`
		Password string `default:"zmemo"`
		Port     string `default:"3306"`
		Host     string `default:"localhost"`
	}
}

// NewConfig config.yamlの設定読み込み
func NewConfig() *Config {
	config := new(Config)

	if err := configor.Load(config, "./config/config.yaml"); err != nil {
		panic(err)
	}
	return config
}
