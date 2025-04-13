package main

import (
	"go_docker_api/api/router"
	"go_docker_api/configs"
)

var logger *configs.Logger

func main() {

	// Get logger
	logger = configs.GetLogger("main")
	logger.Info("Starting the application...")

	// Initialze configs
	logger.Info("Initializing configs...")
	err := configs.Init()

	// Check if there was an error initializing configs
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	router.Initialze()

}
