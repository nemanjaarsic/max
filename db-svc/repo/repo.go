package repo

import (
	"context"
	"max-db-svc/model"
	"max-db-svc/util"
)

type UserRepository interface {
	GetUserByID(context.Context, string) (model.User, error)
	GetByUsername(context.Context, string) (model.User, error)
	GetAllUsers(context.Context) ([]model.User, error)
	GetUserByToken(context.Context, string) (model.User, error)
}

type TransactionRepository interface {
	Insert(context.Context, model.Transaction) error
}

type FinanceRepository interface {
	GetBalanceByUserID(context.Context, string) (float32, util.Error)
	Insert(context.Context, model.Finance) error
	Update(context.Context, model.Finance) error
}

type AuthRepository interface {
	Insert(context.Context, string, string) error
	FindByUserID(context.Context, string) (model.Token, util.Error)
	FindToken(context.Context, string) (model.Token, util.Error)
}

type Repositories struct {
	UserRepo        UserRepository
	TransactionRepo TransactionRepository
	FinanceRepo     FinanceRepository
	AuthRepo        AuthRepository
}

func InitRepos(repo *Repositories) {
	repo.AuthRepo = NewAuthRepo()
	repo.UserRepo = NewUserRepo()
	repo.FinanceRepo = NewFinanceRepo()
	repo.TransactionRepo = NewTransactionRepo()
}
