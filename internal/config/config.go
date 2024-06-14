package config

import (
	"fmt"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"0.0.0.0"`
		Port   string `yaml:"port" env-default:"8080"`
	} `yaml:"listen"`
	Storage        StorageConfig `yaml:"storage"`
	JwtKey         string        `yaml:"jwt_key" env-required:"true"`
	AppVersion     string        `yaml:"app_version" env-required:"true"`
}

type StorageConfig struct {
	Host          string `json:"host"`
	Port          string `json:"port"`
	Database      string `json:"database"`
	Username      string `json:"username"`
	Password      string `json:"password"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {

 once.Do(func() {
		
		pathConfig := "config.yml"

		instance = &Config{}
		if err := cleanenv.ReadConfig(pathConfig, instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			fmt.Println("config.go de errr", help)
		}

	})
	return instance
}
