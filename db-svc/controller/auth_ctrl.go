package controller

import (
	"context"
	"max-db-svc/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

func (c *DatabaseController) SaveToken(ctx context.Context, in *pb.SaveTokenRequest) (*empty.Empty, error) {
	err := c.AuthService.SaveToken(ctx, in.GetUserId(), in.GetToken())

	return &empty.Empty{}, err
}

func (c *DatabaseController) GetTokenByUserID(ctx context.Context, in *pb.GetTokenByUserIdRequest) (*pb.GetTokenResponse, error) {
	t, err := c.AuthService.GetTokenByUserID(ctx, in.GetUserId())

	return &pb.GetTokenResponse{
		Token: t,
	}, err
}

func (c *DatabaseController) GetToken(ctx context.Context, in *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	t, err := c.AuthService.GetToken(ctx, in.GetToken())

	return &pb.GetTokenResponse{
		Token: t,
	}, err
}
