package http

import (
	"context"
	_ "github.com/Enthreeka/go-service-notification/docs"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity/dto"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

type messageHandler struct {
	messageUsecase notification.MessageService

	log *logger.Logger
}

func NewMessageHandler(messageUsecase notification.MessageService, log *logger.Logger) *messageHandler {
	return &messageHandler{
		messageUsecase: messageUsecase,
		log:            log,
	}
}

// GetDetailInfoHandler godoc
// @Summary Get Messages
// @Tags message
// @Description get message
// @Accept json
// @Produce json
// @Param input body dto.IDMessageRequest true "Get detail data from message"
// @Success 200 {object} map[string][]entity.MessageInfo
// @Failure 400 {object} apperror.AppError
// @Failure 500 {object} apperror.AppError
// @Router /api/message/info [get]
func (m *messageHandler) GetDetailInfoHandler(c *fiber.Ctx) error {
	id := &dto.IDMessageRequest{}

	err := c.BodyParser(id)
	if err != nil {
		m.log.Error("failed to parse request body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(apperror.NewError("Invalid request body", err))
	}

	message, err := m.messageUsecase.GetInfoNotification(context.Background(), id.Id)
	if err != nil {
		m.log.Error("get info controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(message)
}

// GetGroupByStatusHandler godoc
// @Summary Get Messages Group By
// @Tags message
// @Description get message
// @Produce json
// @Success 200 {object} map[string][]entity.MessageInfo
// @Failure 500 {object} apperror.AppError
// @Router /api/message/group [get]
func (m *messageHandler) GetGroupByStatusHandler(c *fiber.Ctx) error {
	message, err := m.messageUsecase.GetAllGroupByStatus(context.Background())
	if err != nil {
		m.log.Error("get info group by controller: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(err)
	}

	return c.Status(fiber.StatusOK).JSON(message)
}
