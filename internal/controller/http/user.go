package http

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/usecase"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type clientHandler struct {
	userUsecase usecase.Client

	log *logger.Logger
}

func NewClientHandler(clientUsecase usecase.Client, log *logger.Logger) *clientHandler {
	return &clientHandler{
		userUsecase: clientUsecase,
		log:         log,
	}
}

func (u *clientHandler) CreateUserHandler(c *fiber.Ctx) error {
	client := &entity.Client{}

	err := c.BodyParser(client)
	if err != nil {
		u.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = u.userUsecase.CreateClient(context.Background(), client)
	if err != nil {
		if err == apperror.ErrIncorrectNumber {
			return c.Status(fiber.StatusInternalServerError).JSON(apperror.ErrIncorrectNumber)
		}
		u.log.Error("%v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusCreated).JSON(nil)
}
