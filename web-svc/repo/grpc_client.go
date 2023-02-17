package repo

import (
	"max-web-svc/config"
	"max-web-svc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	DatabaseClientHost    string
	IdempotencyClientHost string
}

func NewGRPCDatabaseClient() (pb.DatabaseClient, error) {
	conn, err := grpc.Dial(config.Conf.DatabaseService.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewDatabaseClient(conn)
	return client, nil
}

func NewGRPCIdempotencyClient() (pb.IdempotencyClient, error) {
	conn, err := grpc.Dial(config.Conf.IdempotencyService.Host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewIdempotencyClient(conn)
	return client, nil
}
