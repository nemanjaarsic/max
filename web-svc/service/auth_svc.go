package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"max-web-svc/model"
	"max-web-svc/repo"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	authRepo repo.AuthRepository
}

func (s *AuthService) Init(repo *repo.Repositories) {
	s.authRepo = repo.AuthRepo
}

func (s *AuthService) Login(ctx context.Context, l model.Login) (string, error) {
	u, err := s.authRepo.GetPassword(ctx, l.Username)
	if err != nil {
		return "", err
	}

	if !checkPasswordHash(u.Password, l.Password) {
		return "", fmt.Errorf("[AuthService Login] Invalid password")
	}

	return u.ID, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, token string) bool {
	t, err := s.authRepo.GetToken(ctx, token)
	if err != nil || t == "" {
		return false
	}
	return true
}

// GetToken return token to client. If token does not exist or is expired new token is created
func (s *AuthService) GetToken(ctx context.Context, userID string) (string, error) {
	t, err := s.authRepo.GetToken(ctx, userID)
	if err != nil {
		return "", err
	}
	if t == "" {
		newToken := generateToken(userID)
		s.authRepo.SaveToken(ctx, userID, newToken)
		return newToken, nil
	}
	return t, nil
}

// checkPasswordHash validate whether provided password matches one in database
func checkPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// generateToken returns a unique hash value, on witch token is based
func generateToken(userID string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(userID), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}
