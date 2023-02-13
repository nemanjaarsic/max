package repo

import (
	"context"
	"max-web-svc/model"
	"max-web-svc/pb"
)

type authRepo struct {
	client pb.DatabaseClient
}

// Check interface compliance
var _ AuthRepository = (*authRepo)(nil)

func NewAuthRepo() *authRepo {
	client, _ := NewGRPCDatabaseClient()
	return &authRepo{
		client: client,
	}
}

func (r *authRepo) GetPassword(ctx context.Context, username string) (model.User, error) {
	u, err := r.client.GetUserByUsername(ctx, &pb.GetUserByUsernameRequest{Username: username})
	if err != nil {
		return model.User{}, err
	}
	return model.User{
		ID:       u.Id,
		Username: u.Username,
		Password: u.Password,
		Name:     u.Name,
	}, nil
}

func (r *authRepo) GetTokenByUserID(ctx context.Context, userID string) (string, error) {
	t, err := r.client.GetTokenByUserId(ctx, &pb.GetTokenByUserIdRequest{UserId: userID})
	if err != nil {
		return "", err
	}
	return t.GetToken(), nil
}

func (r *authRepo) SaveToken(ctx context.Context, userID, token string) error {
	_, err := r.client.SaveToken(ctx, &pb.SaveTokenRequest{UserId: userID, Token: token})
	return err
}

func (r *authRepo) GetToken(ctx context.Context, token string) (string, error) {
	t, err := r.client.GetToken(ctx, &pb.GetTokenRequest{Token: token})
	if err != nil {
		return "", err
	}
	return t.GetToken(), nil
}
