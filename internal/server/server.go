package server

import (
	"context"
	"fmt"
	"github.com/Enthreeka/go-service-notification/internal/config"
	"github.com/Enthreeka/go-service-notification/internal/controller/http"
	postgres2 "github.com/Enthreeka/go-service-notification/internal/repo/postgres"
	"github.com/Enthreeka/go-service-notification/internal/usecase"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"github.com/gofiber/fiber/v2"
)

func Run(log *logger.Logger, cfg *config.Config) error {
	psql, err := postgres.New(context.Background(), cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}
	defer psql.Close()

	clientRepoPG := postgres2.NewClientRepositoryPG(psql)

	clientUsecase := usecase.NewClientUsecase(clientRepoPG, log)

	clientHandler := http.NewClientHandler(clientUsecase, log)

	app := fiber.New()

	app.Post("/create", clientHandler.CreateUserHandler)
	app.Post("/update", clientHandler.UpdateUserHandler)
	app.Delete("/user", clientHandler.DeleteUserHandler)

	log.Info("Starting http server: %s:%s", cfg.HTTPServer.TypeServer, cfg.HTTPServer.Port)
	if err = app.Listen(fmt.Sprintf(":%s", cfg.HTTPServer.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}
