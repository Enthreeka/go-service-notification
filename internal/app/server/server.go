package server

import (
	"context"
	"fmt"
	"github.com/Enthreeka/go-service-notification/internal/config"
	http2 "github.com/Enthreeka/go-service-notification/internal/notification/controller/http"
	pg "github.com/Enthreeka/go-service-notification/internal/notification/repo/postgres"
	"github.com/Enthreeka/go-service-notification/internal/notification/usecase"
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

	clientRepoPG := pg.NewClientRepositoryPG(psql)
	notificationRepoPG := pg.NewNotificationRepositoryPG(psql)
	messageRepoPG := pg.NewMessageRepositoryPG(psql)

	clientUsecase := usecase.NewClientUsecase(clientRepoPG, log)
	notificationUsecase := usecase.NewNotificationUsecase(notificationRepoPG, log)
	messageUsecase := usecase.NewMessageUsecase(messageRepoPG, log)

	clientHandler := http2.NewClientHandler(clientUsecase, log)
	notificationHandler := http2.NewNotificationHandler(notificationUsecase, log)
	messageHandler := http2.NewMessageHandler(messageUsecase, log)

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

	v3 := app.Group("/message")
	v3.Get("/info", messageHandler.GetDetailInfoHandler)
	v3.Get("/group", messageHandler.GetGroupByStatusHandler)

	log.Info("Starting http server: %s:%s", cfg.NotificationHTTTPServer.TypeServer, cfg.NotificationHTTTPServer.Port)
	if err = app.Listen(fmt.Sprintf(":%s", cfg.NotificationHTTTPServer.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}
