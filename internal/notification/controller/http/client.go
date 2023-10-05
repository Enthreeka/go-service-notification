package http

import (
	"context"
	"errors"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type clientHandler struct {
	clientUsecase notification.ClientService

	log *logger.Logger
}

func NewClientHandler(clientUsecase notification.ClientService, log *logger.Logger) *clientHandler {
	return &clientHandler{
		clientUsecase: clientUsecase,
		log:           log,
	}
}

func (u *clientHandler) CreateClientHandler(c *fiber.Ctx) error {
	client := &dto.CreateClientRequest{}

	err := c.BodyParser(client)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = u.clientUsecase.CreateClient(context.Background(), client)
	if err != nil {
		if err == apperror.ErrIncorrectNumber {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrIncorrectNumber)
		}
		u.log.Error("create client controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (u *clientHandler) UpdateClientHandler(c *fiber.Ctx) error {
	client := &dto.UpdateClientRequest{}

	err := c.BodyParser(client)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = u.clientUsecase.UpdateClient(context.Background(), client)
	if err != nil {
		if errors.Is(err, apperror.ErrIncorrectNumber) {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrIncorrectNumber)
		} else if errors.Is(err, apperror.ErrClientAttribute) {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrClientAttribute)
		}

		u.log.Error("update client controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (u *clientHandler) DeleteClientHandler(c *fiber.Ctx) error {
	client := &entity.Client{}

	err := c.BodyParser(client)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = u.clientUsecase.DeleteClient(context.Background(), client.ID)
	if err != nil {
		u.log.Error("delete client controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
