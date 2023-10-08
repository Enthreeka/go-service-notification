package main

import (
	_ "github.com/Enthreeka/go-service-notification/docs"
	"github.com/Enthreeka/go-service-notification/internal/app/server"
	"github.com/Enthreeka/go-service-notification/internal/config"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
)

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger Api for notification service
// @host localhost:8080
// @BasePath /
func main() {
	log := logger.New()

	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed to load config: %v", err)
	}

	if err := server.Run(log, cfg); err != nil {
		log.Fatal("failed to run server: %v", err)
	}
}
