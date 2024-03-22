package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)


type Config struct {
	Env         string     	 	`yaml:"env" env-default:"local"`
	PathEnv 	string 			`yaml:"path_env string" env-default:"config/local.env"`
	Storage     Storage 	 	`yaml:"storage"`
}

type Storage struct {
	Hostname 	string 			`yaml:"hostname" env-required:"true"`
	Port    	string          `yaml:"port"`
	Password 	string 			`yaml:"password" env-default:""`
	Timeout 	time.Duration 	`yaml:"timeout"`
}

func MustLoadPath(configPath string) *Config {
	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("cannot read config: " + err.Error())
	}
	return &cfg
}