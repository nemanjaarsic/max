package controller

import "max-web-svc/service"

type Controllers struct {
	AuthController      *AuthController
	OperationController *OperationController
}

func (c *Controllers) Init(svcs *service.Services) {
	c.AuthController = NewAuthController(svcs)
	c.OperationController = NewOperationController(svcs)
}
