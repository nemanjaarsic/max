package service

import "max-web-svc/repo"

type Services struct {
	AuthService      AuthService
	OperationService OperationService
}

func (svcs *Services) Init(repos *repo.Repositories) {
	svcs.AuthService.Init(repos)
	svcs.OperationService.Init(repos)
}
