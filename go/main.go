package main

import (
	"os"
	"scope3/go-interview/config"

	log "github.com/sirupsen/logrus"

	"scope3/go-interview/cli"
)

func main() {
	config.NewConfig()

	log.SetFormatter(&log.JSONFormatter{
		DisableTimestamp: true,
		PrettyPrint:      true,
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(config.Config.LogLevel)

	cli.Execute()
}
