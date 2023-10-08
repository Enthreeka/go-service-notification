package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/mail"
	"io"
	"net/http"
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

	for _, value := range clientsMessage {
		api := "https://probe.fbrq.cloud/v1/send/" + "123"

		client := &http.Client{}

		body := struct {
			Id    int    `json:"id"`
			Phone string `json:"phone"`
			Text  string `json:"text"`
		}{
			Id:    123,
			Phone: value.PhoneNumber,
			Text:  value.Message,
		}
		bodyByte, _ := json.Marshal(body)

		req, err := http.NewRequest("POST", api, bytes.NewBuffer(bodyByte))
		if err != nil {
			return err
		}

		req.WithContext(ctx)

		req.Header.Set("Authorization", bearer)
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		value.Status = resp.Status

		bodyString := string(bodyBytes)

		fmt.Println(bodyString)

		err = m.mailUsecase.CreateMessageInfo(ctx, &value)
		if err != nil {
			return err
		}
	}

	return nil
}
