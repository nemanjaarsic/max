package service

import (
	"context"
	"fmt"
	"log"
	"max-db-svc/model"
	"max-db-svc/repo"
	"net/http"
)

type FinanceService struct {
	FinanceRepo repo.FinanceRepository
}

func (s *FinanceService) Init(repos *repo.Repositories) {
	s.FinanceRepo = repos.FinanceRepo
}

func (s *FinanceService) Deposit(ctx context.Context, userID string, amount float32) (float32, error) {
	log.Print(ctx.Err())
	b, err := s.FinanceRepo.GetBalanceByUserID(ctx, userID)
	if err.Code != http.StatusMethodNotAllowed && err.Err != nil {
		return -1, err.Err
	}
	//Insert balance for the first time
	if err.Code == http.StatusMethodNotAllowed {
		if err := s.FinanceRepo.Insert(ctx, model.Finance{UserID: userID, Balance: amount}); err != nil {
			return -1, err
		}
		return b, nil
	}

	newBalance := b + amount
	if err := s.FinanceRepo.Update(ctx, model.Finance{UserID: userID, Balance: newBalance}); err != nil {
		return b, err
	}
	log.Print(ctx.Err())
	return newBalance, nil
}

func (s *FinanceService) Withdraw(ctx context.Context, userID string, amount float32) (float32, error) {
	b, err := s.FinanceRepo.GetBalanceByUserID(ctx, userID)
	if err.Code != http.StatusMethodNotAllowed && err.Err != nil {
		return -1, err.Err
	}

	newBalance := b - amount
	if err.Code == http.StatusMethodNotAllowed || newBalance < 0 {
		return -1, fmt.Errorf("[FinanceService Withdraw] Withdraw amount is grater than current balance. Please make a new deposit first")
	}

	if err := s.FinanceRepo.Update(ctx, model.Finance{UserID: userID, Balance: newBalance}); err != nil {
		return b, err
	}
	return newBalance, nil
}

func (s *FinanceService) GetBalance(ctx context.Context, userID string) (float32, error) {
	b, err := s.FinanceRepo.GetBalanceByUserID(ctx, userID)
	if err.Err != nil {
		return -1, err.Err
	}
	return b, nil
}
