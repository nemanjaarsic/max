package service

import (
	"context"
	"max-idempotency-svc/model"
)

type IdempotencyService interface {
	Validate(context.Context, model.Transaction) (string, error)
}

type Services struct {
	IdempotencySvc IdempotencyService
}

func (svcs *Services) InitServices() {
	svcs.IdempotencySvc = NewIdempotencySvc()
}
