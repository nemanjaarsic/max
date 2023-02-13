package repo

import (
	"context"
	"max-web-svc/model"
	"max-web-svc/pb"
)

type operationRepo struct {
	client pb.DatabaseClient
}

// Check interface compliance
var _ OperationRepository = (*operationRepo)(nil)

func NewOperationRepo() *operationRepo {
	client, _ := NewGRPCDatabaseClient()
	return &operationRepo{
		client: client,
	}
}

func (r *operationRepo) Deposit(ctx context.Context, out model.Operation) (float32, error) {
	res, err := r.client.Deposit(ctx, &pb.DepositRequest{
		Id:        out.ID,
		Token:     out.Token,
		Amount:    out.Amount,
		Timestamp: out.Timestamp,
	})
	if err != nil {
		return 0, err
	}

	return res.Balance, err
}

func (r *operationRepo) Withdraw(ctx context.Context, out model.Operation) (float32, error) {
	res, err := r.client.Withdraw(ctx, &pb.WithdrawRequest{
		Id:        out.ID,
		Token:     out.Token,
		Amount:    out.Amount,
		Timestamp: out.Timestamp,
	})
	if err != nil {
		return 0, err
	}
	return res.Balance, err
}

func (r *operationRepo) GetBalance(ctx context.Context, token string) (float32, error) {
	res, err := r.client.GetBalance(ctx, &pb.GetBalanceRequest{
		Token: token,
	})
	if err != nil {
		return -1, err
	}
	return res.Balance, err
}
