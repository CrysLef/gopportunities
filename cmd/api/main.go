package main

import (
	"github.com/CrysLef/gopportunities/config"
	"github.com/CrysLef/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger := config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	// Initialize router
	router.Initialize()
}
