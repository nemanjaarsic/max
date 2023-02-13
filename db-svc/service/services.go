package service

import (
	"max-db-svc/repo"
)

type Services struct {
	FinanceService     FinanceService
	TransactionService TransactionService
	UserService        UserService
	AuthService        AuthService
}

func (svcs *Services) Init(repos *repo.Repositories) {
	svcs.FinanceService.Init(repos)
	svcs.TransactionService.Init(repos)
	svcs.UserService.Init(repos)
	svcs.AuthService.Init(repos)
}
