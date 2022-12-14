package config

import (
	"os"
	"sync"
)

type AppConfig struct {
	SERVER_PORT string
	// DB_DRIVER   string
	DB_HOST     string
	DB_USERNAME string
	DB_PORT     string
	DB_PASSWORD string
	DB_NAME     string
}

var lock = &sync.Mutex{}
var config *AppConfig

func GetConfig() *AppConfig {

	lock.Lock()
	defer lock.Unlock()

	if config == nil {
		config = initConfig()
	}

	return config

}

func initConfig() *AppConfig {

	var defaultConfig AppConfig

	defaultConfig.SERVER_PORT = os.Getenv("SERVER_PORT")
	// defaultConfig.DB_DRIVER = os.Getenv("DB_DRIVER")
	defaultConfig.DB_USERNAME = os.Getenv("DB_USERNAME")
	defaultConfig.DB_HOST = os.Getenv("DB_HOST")
	defaultConfig.DB_NAME = os.Getenv("DB_NAME")
	defaultConfig.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	defaultConfig.DB_PORT = os.Getenv("DB_PORT")

	return &defaultConfig
}
