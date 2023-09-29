package postgres

import (
	"context"
	"fmt"
	"github.com/Enthreeka/go-service-notification/internal/entity"
	"github.com/Enthreeka/go-service-notification/internal/repo"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"strings"
)

type clientRepositoryPG struct {
	*postgres.Postgres
}

func NewClientRepositoryPG(postgres *postgres.Postgres) repo.Client {
	return &clientRepositoryPG{
		postgres,
	}
}

func (c *clientRepositoryPG) Create(ctx context.Context, client *entity.Client) error {
	query := `INSERT INTO client 
		(id,tag,time_zone,phone_number,operator_code)
				VALUES ($1,$2,$3,$4,$5)`

	_, err := c.Pool.Exec(ctx, query,
		client.ID,
		client.Tag,
		client.TimeZone,
		client.PhoneNumber,
		client.OperatorCode,
	)
	return err
}

func (c *clientRepositoryPG) Update(ctx context.Context, client *entity.Client) error {
	var counter int
	args := []string{}
	var builder strings.Builder

	builder.WriteString("UPDATE client SET")

	if client.Tag != "" {
		counter++
		builder.WriteString(fmt.Sprintf(" tag = $%d", counter))
		args = append(args, client.Tag)
	}

	if client.OperatorCode != "" {
		counter++
		builder.WriteString(fmt.Sprintf(" ,operator_code = $%d", counter))
		args = append(args, client.OperatorCode)
	}

	if client.PhoneNumber != "" {
		counter++
		builder.WriteString(fmt.Sprintf(" ,phone_number = $%d", counter))
		args = append(args, client.PhoneNumber)
	}

	if client.TimeZone != "" {
		counter++
		builder.WriteString(fmt.Sprintf(" ,time_zone = $%d", counter))
		args = append(args, client.TimeZone)
	}

	counter++
	builder.WriteString(fmt.Sprintf(" WHERE id = $%d", counter))
	args = append(args, client.ID)

	fmt.Println(builder.String())
	_, err := c.Pool.Exec(ctx, builder.String(), args)
	return err
}

func (c *clientRepositoryPG) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM client WHERE id = $1`

	_, err := c.Pool.Exec(ctx, query, id)
	return err
}
