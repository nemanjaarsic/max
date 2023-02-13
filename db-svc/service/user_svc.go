package service

import (
	"context"
	"log"
	"max-db-svc/model"
	"max-db-svc/repo"
)

type UserService struct {
	UserRepo repo.UserRepository
}

func (s *UserService) Init(repo *repo.Repositories) {
	s.UserRepo = repo.UserRepo
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.UserRepo.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, ID string) (model.User, error) {
	return s.UserRepo.GetUserByID(ctx, ID)
}

func (s *UserService) GetByUsername(ctx context.Context, username string) (model.User, error) {
	return s.UserRepo.GetByUsername(ctx, username)
}

func (s *UserService) GetUserByToken(ctx context.Context, token string) (model.User, error) {
	log.Print(ctx.Err())
	return s.UserRepo.GetUserByToken(ctx, token)
}
