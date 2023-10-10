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
	"github.com/gofiber/swagger"
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
	signalRepoPG := pg.NewSignalRepositoryPG(psql)

	clientUsecase := usecase.NewClientUsecase(clientRepoPG, log)
	notificationUsecase := usecase.NewNotificationUsecase(notificationRepoPG, signalRepoPG, log)
	messageUsecase := usecase.NewMessageUsecase(messageRepoPG, log)

	clientHandler := http2.NewClientHandler(clientUsecase, log)
	notificationHandler := http2.NewNotificationHandler(notificationUsecase, log)
	messageHandler := http2.NewMessageHandler(messageUsecase, log)

	app := fiber.New()

	app.Get("/docs/*", swagger.HandlerDefault)

	v1 := app.Group("/api")

	v2 := v1.Group("/client")
	v2.Post("/create", clientHandler.CreateClientHandler)
	v2.Post("/update", clientHandler.UpdateClientHandler)
	v2.Delete("/delete", clientHandler.DeleteClientHandler)

	v3 := v1.Group("/notification")
	v3.Post("/create", notificationHandler.CreateNotificationHandler)
	v3.Post("/update", notificationHandler.UpdateNotificationHandler)
	v3.Post("/stat", notificationHandler.GetStatNotificationHandler)
	v3.Delete("/delete", notificationHandler.DeleteNotificationHandler)

	v4 := v1.Group("/message")
	v4.Get("/info/:id", messageHandler.GetDetailInfoHandler)
	v4.Get("/group", messageHandler.GetGroupByStatusHandler)

	log.Info("Starting http server: %s:%s", cfg.NotificationHTTTPServer.TypeServer, cfg.NotificationHTTTPServer.Port)
	if err = app.Listen(fmt.Sprintf(":%s", cfg.NotificationHTTTPServer.Port)); err != nil {
		log.Fatal("Server listening failed:%s", err)
	}

	return nil
}
