package controller

import (
	"context"
	"max-db-svc/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

func (c *DatabaseController) GetAllUsers(ctx context.Context, empty *empty.Empty) (*pb.UsersResponse, error) {
	users, err := c.UserService.GetAllUsers(ctx)
	resp := make([]*pb.UserResponse, 0)

	for _, u := range users {
		resp = append(resp, &pb.UserResponse{
			Id:       u.ID,
			Username: u.Username,
			Name:     u.Name,
		})
	}

	return &pb.UsersResponse{
		Users: resp,
	}, err
}

func (c *DatabaseController) GetUserByID(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	user, err := c.UserService.GetUserByID(ctx, in.GetId())

	return &pb.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Name:     user.Name,
	}, err
}

func (c *DatabaseController) GetUserByUsername(ctx context.Context, in *pb.GetUserByUsernameRequest) (*pb.UserResponse, error) {
	user, err := c.UserService.GetByUsername(ctx, in.GetUsername())

	return &pb.UserResponse{
		Id:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Name:     user.Name,
	}, err
}
