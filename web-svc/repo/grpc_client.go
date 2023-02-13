package repo

import (
	"max-web-svc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	DatabaseClientHost    string
	IdempotencyClientHost string
}

func NewGRPCDatabaseClient() (pb.DatabaseClient, error) {
	conn, err := grpc.Dial("0.0.0.0:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewDatabaseClient(conn)
	return client, nil
}

func NewGRPCIdempotencyClient() (pb.IdempotencyClient, error) {
	conn, err := grpc.Dial("0.0.0.0:7000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := pb.NewIdempotencyClient(conn)
	return client, nil
}
