package repo

import (
	"context"
	"max-web-svc/model"
	"max-web-svc/pb"
)

type idempotencyRepo struct {
	client pb.IdempotencyClient
}

// Check interface compliance
var _ IdempotencyRepository = (*idempotencyRepo)(nil)

func NewIdempotencyRepo() *idempotencyRepo {
	client, _ := NewGRPCIdempotencyClient()
	return &idempotencyRepo{
		client: client,
	}
}

func (r *idempotencyRepo) ValidateTransaction(ctx context.Context, tx model.IdempotencyTx) (string, error) {
	val, err := r.client.ValidateRequest(ctx, &pb.ValidateTransactionRequest{
		Id:        tx.ID,
		Amount:    tx.Amount,
		Timestamp: tx.Timestamp,
	})

	return val.GetValue(), err
}
