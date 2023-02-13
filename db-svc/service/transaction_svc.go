package service

import (
	"context"
	"fmt"
	"log"
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
	log.Print(ctx.Err())
	if tx.OperationType == model.Withdraw {
		log.Print(ctx.Err())
		b, err := s.FinanceRepo.GetBalanceByUserID(ctx, tx.UserID)
		if err.Err != nil {
			return err.Err
		} else if b-tx.Amount < 0 {
			return fmt.Errorf("[TransactionService Insert] Withdraw amount is larger than current balance")
		}
	}
	log.Print(ctx.Err())
	return s.TransactionRepo.Insert(ctx, tx)
}
