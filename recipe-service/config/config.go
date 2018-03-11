// Package config parse config from config.json
package config

import (
	"encoding/json"
	"github.com/hellofreshdevtests/sahilahmadlone-api-test/recipe-service/logger"
	"os"
	"sync"
)

type DBConfig struct {
	Server   []string `json:"server"`
	DBName   string   `json:"dbname"`
	Timeout  int      `json:"timeout"`
	UserName string   `json:"username"`
	Password string   `json:"password"`
}

type AuthConfig struct {
	Type     string `json:"type"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type Config struct {
	ReadTimeout  int `json:"read_timeout"`
	WriteTimeout int `json:"write_timeout"`
	DBConfig     `json:"db"`
	AuthConfig   `json:"auth"`
}

const CONFIG_ARG = "config"

var config Config
var once sync.Once

/*
 * Initializes config object based on path argument otherwise the default
 * This can be used for injecting different configs for different environments
 * If there is error in opening config its a fatal error
 */
func GetConfig() Config {
	once.Do(func() {

		configPath := os.Getenv("CONFIG")
		if configPath == "" {
			logger.Fatal("Error:Config environment variable not found")
		}
		logger.Info("Using config file at %v", configPath)
		configFile, err := os.Open(configPath)
		if err != nil {
			logger.Fatal("Error:Config file Path", err)
		}
		defer configFile.Close()
		jsonParser := json.NewDecoder(configFile)
		if err = jsonParser.Decode(&config); err == nil {
			logger.Info("successfully parsed config %+v", config)
		} else {
			logger.Fatal("Error:", err)
		}
	})

	return config
}
