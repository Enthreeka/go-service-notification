package usecase

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
	"github.com/Enthreeka/go-service-notification/internal/repo"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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

func (c *clientUsecase) CreateClient(ctx context.Context, request *dto.CreateClientRequest) error {
	if !entity.IsCorrectNumber(request.PhoneNumber) {
		return apperror.ErrIncorrectNumber
	}

	location, err := time.LoadLocation(request.TimeZone)
	if err != nil {
		return apperror.NewError("failed to load time zone", err)
	}

	client := &entity.Client{
		PhoneNumber:      request.PhoneNumber,
		ClientProperty:   request.ClientProperty,
		TimeZone:         time.Now().In(location),
		ClientPropertyID: uuid.New().String(),
	}

	err = c.clientRepoPG.Create(ctx, client)
	if err != nil {
		return apperror.NewError("failed to create client in postgres", err)
	}

	return nil
}

func (c *clientUsecase) UpdateClient(ctx context.Context, request *dto.UpdateClientRequest) error {
	if request.PhoneNumber != "" {
		if !entity.IsCorrectNumber(request.PhoneNumber) {
			return apperror.ErrIncorrectNumber
		}
	}
	client := &entity.Client{
		ID:               request.ID,
		ClientPropertyID: request.ClientPropertyID,
		PhoneNumber:      request.PhoneNumber,
		ClientProperty:   request.ClientProperty,
	}

	if request.TimeZone != "" {
		location, err := time.LoadLocation(request.TimeZone)
		if err != nil {
			return apperror.NewError("failed to load time zone", err)
		}
		client.TimeZone = time.Now().In(location)
	}

	err := c.clientRepoPG.Update(ctx, client)
	if err != nil {
		if err == pgx.ErrNoRows {
			return apperror.ErrClientAttribute
		}
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
