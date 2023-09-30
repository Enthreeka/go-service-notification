package http

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
	"github.com/Enthreeka/go-service-notification/internal/usecase"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type notificationHandler struct {
	notificationUsecase usecase.Notification

	log *logger.Logger
}

func NewNotificationHandler(notificationUsecase usecase.Notification, log *logger.Logger) *notificationHandler {
	return &notificationHandler{
		notificationUsecase: notificationUsecase,
		log:                 log,
	}
}

func (n *notificationHandler) CreateNotificationHandler(c *fiber.Ctx) error {
	notificationRequest := &entity.Notification{}

	err := c.BodyParser(notificationRequest)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = n.notificationUsecase.CreateNotification(context.Background(), notificationRequest)
	if err != nil {
		if err == apperror.ErrIncorrectTime {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrIncorrectTime)
		}
		n.log.Error("create notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (n *notificationHandler) UpdateNotificationHandler(c *fiber.Ctx) error {
	notificationRequest := &dto.UpdateNotificationRequest{}

	err := c.BodyParser(notificationRequest)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = n.notificationUsecase.UpdateNotification(context.Background(), notificationRequest)
	if err != nil {
		if err == apperror.ErrIncorrectTime {
			return c.Status(fiber.StatusBadRequest).JSON(apperror.ErrIncorrectTime)
		}
		n.log.Error("update notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (n *notificationHandler) DeleteNotificationHandler(c *fiber.Ctx) error {
	notificationRequest := &dto.DeleteNotificationRequest{}

	err := c.BodyParser(notificationRequest)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	err = n.notificationUsecase.DeleteNotification(context.Background(), notificationRequest)
	if err != nil {
		n.log.Error("delete notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (n *notificationHandler) GetStatNotificationHandler(c *fiber.Ctx) error {
	notificationRequest := &dto.GetNotificationRequest{}

	err := c.BodyParser(notificationRequest)
	if err != nil {
		n.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	stat, err := n.notificationUsecase.GetByCreateAt(context.Background(), notificationRequest)
	if err != nil {
		n.log.Error("get notification controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(stat)
}
