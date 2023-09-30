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
	notificationRepoPG := postgres2.NewNotificationRepositoryPG(psql)

	clientUsecase := usecase.NewClientUsecase(clientRepoPG, log)
	notificationUsecase := usecase.NewNotificationUsecase(notificationRepoPG, log)

	clientHandler := http.NewClientHandler(clientUsecase, log)
	notificationHandler := http.NewNotificationHandler(notificationUsecase, log)

	app := fiber.New()

	v1 := app.Group("/client")
	v1.Post("/create", clientHandler.CreateClientHandler)
	v1.Post("/update", clientHandler.UpdateClientHandler)
	v1.Delete("/delete", clientHandler.DeleteClientHandler)

	v2 := app.Group("/notification")
	v2.Post("/create", notificationHandler.CreateNotificationHandler)
	v2.Delete("/:id", notificationHandler.DeleteNotificationHandler)

	log.Info("Starting http server: %s:%s", cfg.HTTPServer.TypeServer, cfg.HTTPServer.Port)
	if err = app.Listen(fmt.Sprintf(":%s", cfg.HTTPServer.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}
