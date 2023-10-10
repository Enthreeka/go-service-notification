package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Enthreeka/go-service-notification/internal/apperror"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/mail"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"strconv"
)

type mailRequest struct {
	mailUsecase mail.MailService

	signalMessageCh chan []entity.ClientsMessage
}

func NewMailRequest(mailUsecase mail.MailService, signalMessageCh chan []entity.ClientsMessage) *mailRequest {
	return &mailRequest{
		mailUsecase:     mailUsecase,
		signalMessageCh: signalMessageCh,
	}
}

func (m *mailRequest) SendRequestAPI(ctx context.Context, token string, clientsMessage []entity.ClientsMessage) error {
	bearer := "Bearer " + token

	id := uuid.New().String()

	for _, value := range clientsMessage {
		value.ID = id

		select {
		case <-ctx.Done():
			err := m.mailUsecase.CreateMessageInfo(context.Background(), &value)
			if err != nil {
				return err
			}

			return apperror.NewError("time has disappeared", ctx.Err())
		default:
			randID := rand.Int()

			api := fmt.Sprintf("https://probe.fbrq.cloud/v1/send/%d", randID)

			client := &http.Client{}

			body := struct {
				Id    int    `json:"id"`
				Phone string `json:"phone"`
				Text  string `json:"text"`
			}{
				Id:    randID,
				Phone: value.PhoneNumber,
				Text:  value.Message,
			}
			bodyByte, _ := json.Marshal(body)

			req, err := http.NewRequest("POST", api, bytes.NewBuffer(bodyByte))
			if err != nil {
				return err
			}

			req.Header.Set("Authorization", bearer)
			req.Header.Add("Accept", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				return err
			}

			resp.Body.Close()

			//bodyBytes, err := io.ReadAll(resp.Body)
			//if err != nil {
			//	return err
			//}

			value.InTime = true
			value.Status = strconv.Itoa(resp.StatusCode)

			err = m.mailUsecase.CreateMessageInfo(context.Background(), &value)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *mailRequest) SendRequestAPIAfterSignal(ctx context.Context, token string, clientsMessage []entity.ClientsMessage) error {
	bearer := "Bearer " + token

	id := uuid.New().String()
	for _, value := range clientsMessage {
		randID := rand.Int()
		value.ID = id

		api := fmt.Sprintf("https://probe.fbrq.cloud/v1/send/%d", randID)

		client := &http.Client{}

		body := struct {
			Id    int    `json:"id"`
			Phone string `json:"phone"`
			Text  string `json:"text"`
		}{
			Id:    randID,
			Phone: value.PhoneNumber,
			Text:  value.Message,
		}
		bodyByte, _ := json.Marshal(body)

		req, err := http.NewRequest("POST", api, bytes.NewBuffer(bodyByte))
		if err != nil {
			return err
		}

		req.Header.Set("Authorization", bearer)
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		resp.Body.Close()

		//bodyBytes, err := io.ReadAll(resp.Body)
		//if err != nil {
		//	return err
		//}

		value.InTime = true
		value.Status = strconv.Itoa(resp.StatusCode)

		err = m.mailUsecase.CreateMessageInfo(context.Background(), &value)
		if err != nil {
			return err
		}
	}

	return nil
}
