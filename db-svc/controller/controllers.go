package controller

import (
	"max-db-svc/pb"
	"max-db-svc/service"
)

type DatabaseController struct {
	pb.UnimplementedDatabaseServer
	FinanceService     *service.FinanceService
	TransactionService *service.TransactionService
	UserService        *service.UserService
	AuthService        *service.AuthService
}

func Init(svcs *service.Services) *DatabaseController {
	ctrl := &DatabaseController{}
	ctrl.FinanceService = &svcs.FinanceService
	ctrl.TransactionService = &svcs.TransactionService
	ctrl.UserService = &svcs.UserService
	ctrl.AuthService = &svcs.AuthService

	return ctrl
}
