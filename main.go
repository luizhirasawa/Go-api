package main

import (
	"github.com/luizhirasawa/Go-api.git/config"
	"github.com/luizhirasawa/Go-api.git/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("Go-API: ")
	err := config.Init()
	if err != nil {
		logger.ErrorF("Error initializing database: %v", err)
		return
	}

	router.Initialize()
}
