package controller

import (
	"context"
	"max-idempotency-svc/model"
	"max-idempotency-svc/pb"
)

func (c *IdempotencyController) ValidateRequest(ctx context.Context, in *pb.ValidateTransactionRequest) (*pb.ValidateTransactionRespons, error) {
	val, err := c.IdempotencyService.Validate(ctx, model.Transaction{
		ID:        in.GetId(),
		Amount:    in.GetAmount(),
		Timestamp: in.GetTimestamp(),
	})
	if err != nil {
		return &pb.ValidateTransactionRespons{}, err
	}

	return &pb.ValidateTransactionRespons{
		Value: val,
	}, err
}
