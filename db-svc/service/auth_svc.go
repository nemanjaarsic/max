package service

import (
	"context"
	"max-db-svc/repo"
	"net/http"
)

type AuthService struct {
	AuthRepo repo.AuthRepository
}

func (s *AuthService) Init(repo *repo.Repositories) {
	s.AuthRepo = repo.AuthRepo
}

func (s *AuthService) SaveToken(ctx context.Context, userID, token string) error {
	return s.AuthRepo.Insert(ctx, userID, token)
}

func (s *AuthService) GetTokenByUserID(ctx context.Context, userID string) (string, error) {
	t, err := s.AuthRepo.FindByUserID(ctx, userID)
	if err.Err != nil {
		return "", err.Err
	}
	if !t.Active || err.Code == http.StatusNotFound {
		return "", nil
	}
	return t.Token, nil
}

func (s *AuthService) GetToken(ctx context.Context, token string) (string, error) {
	t, err := s.AuthRepo.FindToken(ctx, token)
	if err.Err != nil {
		return "", err.Err
	}
	if !t.Active || err.Code == http.StatusNotFound {
		return "", nil
	}
	return t.Token, nil
}
