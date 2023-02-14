package controller

import (
	"context"
	"max-db-svc/model"
	"max-db-svc/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *DatabaseController) Deposit(ctx context.Context, in *pb.DepositRequest) (*pb.OperationRespons, error) {
	user, err := c.UserService.GetUserByToken(ctx, in.GetToken())
	if err != nil {
		return &pb.OperationRespons{}, status.Errorf(
			codes.Internal,
			err.Error(),
		)
	}
	err = c.TransactionService.Insert(ctx, model.Transaction{
		ID:            in.GetId(),
		UserID:        user.ID,
		OperationType: model.Deposit,
		Amount:        in.GetAmount(),
		Timestamp:     in.GetTimestamp(),
	})
	if err != nil {
		return &pb.OperationRespons{}, status.Errorf(
			codes.Internal,
			err.Error(),
		)
	}

	balance, err := c.FinanceService.Deposit(ctx, user.ID, in.GetAmount())

	return &pb.OperationRespons{
		Balance: balance,
		Error:   0,
	}, err
}

func (c *DatabaseController) Withdraw(ctx context.Context, in *pb.WithdrawRequest) (*pb.OperationRespons, error) {
	user, err := c.UserService.GetUserByToken(ctx, in.GetToken())
	if err != nil {
		return &pb.OperationRespons{}, status.Errorf(
			codes.Internal,
			err.Error(),
		)
	}
	err = c.TransactionService.Insert(ctx, model.Transaction{
		ID:            in.GetId(),
		UserID:        user.ID,
		OperationType: model.Withdraw,
		Amount:        in.GetAmount(),
		Timestamp:     in.GetTimestamp(),
	})
	if err != nil {
		return &pb.OperationRespons{}, status.Errorf(
			codes.InvalidArgument,
			err.Error(),
		)
	}

	balance, err := c.FinanceService.Withdraw(ctx, user.ID, in.GetAmount())

	return &pb.OperationRespons{
		Balance: balance,
		Error:   0,
	}, err
}

func (c *DatabaseController) GetBalance(ctx context.Context, in *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	user, err := c.UserService.GetUserByToken(ctx, in.GetToken())
	if err != nil {
		return &pb.GetBalanceResponse{}, status.Errorf(
			codes.Internal,
			err.Error(),
		)
	}

	balance, err := c.FinanceService.GetBalance(ctx, user.ID)

	return &pb.GetBalanceResponse{Balance: balance}, err
}
