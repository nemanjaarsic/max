package service

import (
	"context"
	"fmt"
	"max-db-svc/model"
	"max-db-svc/repo"
)

type TransactionService struct {
	TransactionRepo repo.TransactionRepository
	FinanceRepo     repo.FinanceRepository
	UserRepo        repo.UserRepository
}

func (s *TransactionService) Init(repo *repo.Repositories) {
	s.TransactionRepo = repo.TransactionRepo
	s.FinanceRepo = repo.FinanceRepo
	s.UserRepo = repo.UserRepo
}

func (s *TransactionService) Insert(ctx context.Context, tx model.Transaction) error {
	if tx.OperationType == model.Withdraw {
		b, err := s.FinanceRepo.GetBalanceByUserID(ctx, tx.UserID)
		if err.Err != nil {
			return err.Err
		} else if b-tx.Amount < 0 {
			return fmt.Errorf("[TransactionService Insert] Withdraw amount is larger than current balance")
		}
	}
	return s.TransactionRepo.Insert(ctx, tx)
}
