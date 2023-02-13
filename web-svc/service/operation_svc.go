package service

import (
	"context"
	"max-web-svc/model"
	"max-web-svc/repo"
)

type OperationService struct {
	operationRepo   repo.OperationRepository
	idempotencyRepo repo.IdempotencyRepository
}

func (s *OperationService) Init(repo *repo.Repositories) {
	s.operationRepo = repo.OperationRepo
	s.idempotencyRepo = repo.IdempotencyRepo
}

func (s *OperationService) Deposit(ctx context.Context, deposit model.Operation) (float32, error) {
	var b float32
	var err error
	val, err := s.idempotencyRepo.ValidateTransaction(ctx, model.IdempotencyTx{
		ID:        deposit.ID,
		Amount:    deposit.Amount,
		Timestamp: deposit.Timestamp,
	})
	//Transaction id exists in redis db or error occured while attempting to retrive data
	if val != "" || err != nil {
		b, err = s.operationRepo.GetBalance(ctx, deposit.Token)
	} else {
		b, err = s.operationRepo.Deposit(ctx, deposit)
	}
	return b, err
}

func (s *OperationService) Withdraw(ctx context.Context, withdraw model.Operation) (float32, error) {
	var b float32
	var err error
	val, err := s.idempotencyRepo.ValidateTransaction(ctx, model.IdempotencyTx{
		ID:        withdraw.ID,
		Amount:    withdraw.Amount,
		Timestamp: withdraw.Timestamp,
	})
	//Transaction id exists in redis db or error occured while attempting to retrive data
	if val != "" || err != nil {
		b, err = s.operationRepo.GetBalance(ctx, withdraw.Token)
	} else {
		b, err = s.operationRepo.Withdraw(ctx, withdraw)
	}
	return b, err
}
