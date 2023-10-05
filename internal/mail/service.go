package mail

import "context"

type MailService interface {
	GetTime(ctx context.Context) error
}
