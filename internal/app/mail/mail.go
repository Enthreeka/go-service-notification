package mail

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/config"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/mail/controller/http"
	"github.com/Enthreeka/go-service-notification/internal/mail/repo"
	"github.com/Enthreeka/go-service-notification/internal/mail/usecase"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"sync"
	"time"
)

func Run(log *logger.Logger, cfg *config.Config) error {
	psql, err := postgres.New(context.Background(), cfg.Postgres.URL)
	if err != nil {
		log.Fatal("failed to connect PostgreSQL: %v", err)
	}
	defer psql.Close()

	signalMessageCh := make(chan []entity.ClientsMessage)

	mailRepo := repo.NewMailRepositoryPG(psql)
	mailUsecase := usecase.NewMailUsecase(mailRepo, log)
	mailRequest := http.NewMailRequest(mailUsecase, signalMessageCh)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case clientsMessage := <-signalMessageCh:

				deadline := clientsMessage[0].ExpiresAt
				parsedDeadline, err := time.Parse("2006-01-02 15:04:05 -0700 UTC", deadline)
				if err != nil {
					log.Error("%v", err)
				}

				ctx, cancel := context.WithDeadline(context.Background(), parsedDeadline)

				err = mailRequest.SendRequestAPI(ctx, cfg.ExternalAPI.JWT, clientsMessage)
				if err != nil {
					log.Error("%v", err)
				}

				log.Info("%v", clientsMessage)
				cancel()
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		timeChecker := time.NewTicker(1 * time.Minute)
		for {
			select {
			case <-timeChecker.C:
				clientsMessage, err := mailUsecase.GetTime(context.Background())
				if err != nil {
					log.Error("%v", err)
				}

				if len(clientsMessage) == 0 {
					log.Info("no one message in: %v", time.Now())
				} else {
					log.Info("there is a new notification")
					signalMessageCh <- clientsMessage
				}

			}

		}
	}()

	wg.Wait()

	return nil
}
