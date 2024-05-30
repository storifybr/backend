package main

import (
	"github.com/storify/backend/configs"
	"github.com/storify/backend/internal/router"
)

var (
	logger *configs.Logger
)

func main () {
	logger = configs.GetLogger("main")

	err := configs.InitializeDatabase()
	if err != nil {
		logger.Errorf("database initialization failed: %v", err)
		return
	}

	router.Initialize()
}
