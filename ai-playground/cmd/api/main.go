package main

import (
	"github.com/arymaulanamalik/ai-engineer-playground/ai-playground/configs"
	"github.com/arymaulanamalik/ai-engineer-playground/ai-playground/pkg/logger"
)

func main() {
	config := configs.NewConfig()

	// Setup logger
	logger.Setup(config)

	// Initialize dependencies
	setupDependencies(config)
}

func setupDependencies(config *configs.Config) {

}
