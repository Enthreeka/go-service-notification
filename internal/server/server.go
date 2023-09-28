package server

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/config"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
)

func Run(log *logger.Logger, cfg *config.Config) error {

	psql, err := postgres.New(context.Background(), cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}
	defer psql.Close()

	return nil
}
