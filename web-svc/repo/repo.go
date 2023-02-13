package repo

import (
	"context"
	"max-web-svc/model"
)

type AuthRepository interface {
	GetPassword(context.Context, string) (model.User, error)
	GetToken(context.Context, string) (string, error)
	SaveToken(context.Context, string, string) error
}

type OperationRepository interface {
	Deposit(context.Context, model.Operation) (float32, error)
	Withdraw(context.Context, model.Operation) (float32, error)
	GetBalance(context.Context, string) (float32, error)
}

type IdempotencyRepository interface {
	ValidateTransaction(context.Context, model.IdempotencyTx) (string, error)
}

type Repositories struct {
	AuthRepo        AuthRepository
	OperationRepo   OperationRepository
	IdempotencyRepo IdempotencyRepository
}

func InitRepos(repo *Repositories) {
	repo.AuthRepo = NewAuthRepo()
	repo.OperationRepo = NewOperationRepo()
	repo.IdempotencyRepo = NewIdempotencyRepo()
}
