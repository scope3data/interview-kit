package config

import (
	"os"

	"github.com/joho/godotenv"

	log "github.com/sirupsen/logrus"
)

type ConfigStruct struct {
	ApiKey string
	LogLevel log.Level
}

var Config *ConfigStruct
func NewConfig() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatal(err)
	}

	config := &ConfigStruct{
		ApiKey: os.Getenv("API_KEY"),
		LogLevel: level,
	}

	if config.ApiKey == "" {
		panic("API_KEY environment variable not set")
	}

	Config = config
}
