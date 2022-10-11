package main

import (
	"flag"
	"sport_helper/internal/config"
	"sport_helper/pkg/logger"
)

func main() {
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	if err := config.LoadData(*configPath); err != nil {
		panic(err)
	}

	logger.Setup(*config.GetConfig())

	app := NewApp()
	app.Start()
}
