package postgres

import (
	"context"
	"github.com/Enthreeka/go-service-notification/internal/notification"
	"github.com/Enthreeka/go-service-notification/pkg/postgres"
	"time"
)

type signalRepositoryPG struct {
	*postgres.Postgres
}

func NewSignalRepositoryPG(postgres *postgres.Postgres) notification.Signal {
	return &signalRepositoryPG{
		postgres,
	}
}

func (s *signalRepositoryPG) Create(ctx context.Context, t string) error {
	query := `INSERT INTO signal (created_at) VALUES ($1)`

	createdAtTime, err := time.Parse("15:04 02.01.2006", t)
	if err != nil {
		return err
	}

	_, err = s.Pool.Exec(ctx, query, createdAtTime)
	if err != nil {
		return err
	}

	return nil
}
