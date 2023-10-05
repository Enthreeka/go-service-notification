package postgres

import (
	"context"
	"fmt"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"strings"
)

type clientRepositoryPG struct {
	*postgres.Postgres
}

func NewClientRepositoryPG(postgres *postgres.Postgres) notification.ClientStorage {
	return &clientRepositoryPG{
		Postgres: postgres,
	}
}

func (c *clientRepositoryPG) getClientPropertiesID(ctx context.Context, client *entity.ClientProperty) (string, error) {
	query := `SELECT id FROM client_properties WHERE tag = $1 AND operator_code = $2`
	var id string

	err := c.Pool.QueryRow(ctx, query, client.Tag, client.OperatorCode).Scan(&id)
	return id, err
}

func (c *clientRepositoryPG) Create(ctx context.Context, client *entity.Client) error {
	queryClient := `INSERT INTO client 
		(id_client_properties,time_zone,phone_number)
				VALUES ($1,$2,$3)`

	queryClientProperty := `INSERT INTO client_properties
							(id,tag,operator_code)
							VALUES ($1,$2,$3)`

	tx, err := c.Pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(context.TODO())
		} else {
			tx.Commit(context.TODO())
		}
	}()

	clientPropertyID, err := c.getClientPropertiesID(ctx, &client.ClientProperty)
	if err != nil {
		if err == pgx.ErrNoRows {
			_, err = tx.Exec(ctx, queryClientProperty,
				client.ClientPropertyID,
				client.ClientProperty.Tag,
				client.ClientProperty.OperatorCode,
			)
		} else {
			return err
		}
	} else {
		client.ClientPropertyID = clientPropertyID
	}

	_, err = tx.Exec(ctx, queryClient,
		client.ClientPropertyID,
		client.TimeZone,
		client.PhoneNumber,
	)

	return err
}

// Необходимо реализовать обновление атрибутов клиента
// Если попало поле тега с оператором кода, сперва проверяем, есть ли такие в таблице client_properties
// В случае, если есть получаем id из данной таблицы, и обновляем с новыми атрибутами и новым внешним ключом
// В случае, если нет создаем новое поле в таблице client_properties и добавляем его к новым атрибутом

func (c *clientRepositoryPG) Update(ctx context.Context, client *entity.Client) error {
	var counter int
	args := []interface{}{}
	var builder strings.Builder

	builder.WriteString("UPDATE client SET")

	if client.ClientProperty.Tag != "" || client.ClientProperty.OperatorCode != "" {
		clientPropertyID, err := c.getClientPropertiesID(ctx, &client.ClientProperty)
		if err != nil {
			if err == pgx.ErrNoRows {
				return pgx.ErrNoRows
			} else {
				return err
			}
		}

		counter++
		builder.WriteString(fmt.Sprintf(" id_client_properties = $%d", counter))
		args = append(args, clientPropertyID)
	}

	if client.PhoneNumber != "" {
		counter++
		if counter == 1 {
			builder.WriteString(fmt.Sprintf(" phone_number = $%d", counter))
		} else {
			builder.WriteString(fmt.Sprintf(" ,phone_number = $%d", counter))
		}
		args = append(args, client.PhoneNumber)
	}

	if !client.TimeZone.IsZero() {
		counter++
		if counter == 1 {
			builder.WriteString(fmt.Sprintf(" time_zone = $%d", counter))
		} else {
			builder.WriteString(fmt.Sprintf(" ,time_zone = $%d", counter))
		}
		args = append(args, client.TimeZone)
	}

	counter++
	builder.WriteString(fmt.Sprintf(" WHERE id = $%d", counter))
	args = append(args, client.ID)

	_, err := c.Pool.Exec(ctx, builder.String(), args...)
	return err
}

func (c *clientRepositoryPG) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM client WHERE id = $1`

	_, err := c.Pool.Exec(ctx, query, id)
	return err
}
