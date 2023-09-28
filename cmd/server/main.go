package main

import (
	"github.com/Enthreeka/go-service-notification/internal/config"
	"github.com/Enthreeka/go-service-notification/internal/server"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
)

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
