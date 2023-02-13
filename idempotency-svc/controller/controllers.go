package controller

import (
	"max-idempotency-svc/pb"
	"max-idempotency-svc/service"
)

type IdempotencyController struct {
	pb.UnimplementedIdempotencyServer
	IdempotencyService service.IdempotencyService
}

func Init(svcs *service.Services) *IdempotencyController {
	ctrl := &IdempotencyController{}
	ctrl.IdempotencyService = svcs.IdempotencySvc

	return ctrl
}
