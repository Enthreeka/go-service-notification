package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/repo"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/google/uuid"
	"time"
)

type clientUsecase struct {
	clientRepoPG repo.Client

	log *logger.Logger
}

func NewClientUsecase(clientRepoPG repo.Client, log *logger.Logger) Client {
	return &clientUsecase{
		clientRepoPG: clientRepoPG,
		log:          log,
	}
}

func (c *clientUsecase) CreateClient(ctx context.Context, client *entity.Client) error {
	if !entity.IsCorrectNumber(client.PhoneNumber) {
		return apperror.ErrIncorrectNumber
	}

	client.ID = uuid.New().String()
	location, err := time.LoadLocation(client.TimeZoneDTO)
	if err != nil {
		return apperror.NewError("failed to load time zone", err)
	}

	client.TimeZone = time.Now().In(location)

	err = c.clientRepoPG.Create(ctx, client)
	if err != nil {
		return apperror.NewError("failed to create client in postgres", err)
	}

	return nil
}

func (c *clientUsecase) UpdateClient(ctx context.Context, client *entity.Client) error {
	if client.PhoneNumber != "" {
		if !entity.IsCorrectNumber(client.PhoneNumber) {
			return apperror.ErrIncorrectNumber
		}
	}

	err := c.clientRepoPG.Update(ctx, client)
	if err != nil {
		return apperror.NewError("failed to update client in postgres", err)
	}

	return nil
}

func (c *clientUsecase) DeleteClient(ctx context.Context, id string) error {
	err := c.clientRepoPG.Delete(ctx, id)
	if err != nil {
		return apperror.NewError("failed to delete client from postgres", err)
	}

	return nil
}
