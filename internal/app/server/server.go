package server

import (
	"context"
	"fmt"
	"github.com/Enthreeka/go-service-notification/internal/config"
	http2 "github.com/Enthreeka/go-service-notification/internal/notification/controller/http"
	postgres3 "github.com/Enthreeka/go-service-notification/internal/notification/repo/postgres"
	usecase2 "github.com/Enthreeka/go-service-notification/internal/notification/usecase"
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

	clientRepoPG := postgres3.NewClientRepositoryPG(psql)
	notificationRepoPG := postgres3.NewNotificationRepositoryPG(psql)

	clientUsecase := usecase2.NewClientUsecase(clientRepoPG, log)
	notificationUsecase := usecase2.NewNotificationUsecase(notificationRepoPG, log)

	clientHandler := http2.NewClientHandler(clientUsecase, log)
	notificationHandler := http2.NewNotificationHandler(notificationUsecase, log)

	app := fiber.New()

	v1 := app.Group("/client")
	v1.Post("/create", clientHandler.CreateClientHandler)
	v1.Post("/update", clientHandler.UpdateClientHandler)
	v1.Delete("/delete", clientHandler.DeleteClientHandler)

	v2 := app.Group("/notification")
	v2.Post("/create", notificationHandler.CreateNotificationHandler)
	v2.Post("/update", notificationHandler.UpdateNotificationHandler)
	v2.Post("/stata", notificationHandler.GetStatNotificationHandler)
	v2.Delete("/delete", notificationHandler.DeleteNotificationHandler)

	log.Info("Starting http server: %s:%s", cfg.NotificationHTTTPServer.TypeServer, cfg.NotificationHTTTPServer.Port)
	if err = app.Listen(fmt.Sprintf(":%s", cfg.NotificationHTTTPServer.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}
