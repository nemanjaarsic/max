package main

import (
	"log"
	"max-idempotency-svc/controller"
	"max-idempotency-svc/pb"
	"max-idempotency-svc/service"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	svcs service.Services
)

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:7000")
	log.Print(lis.Addr())
	if err != nil {
		log.Fatalf("Failed to listen on port 7000: %v", err)
	}

	initServices()

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterIdempotencyServer(grpcServer, controller.Init(&svcs))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 7000: %v", err)
	}
}

func initServices() {
	svcs.InitServices()
}
