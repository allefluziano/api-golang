package main

import (
	"github.com/allefluziano/api-golang/config"
	"github.com/allefluziano/api-golang/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
