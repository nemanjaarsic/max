package main

import (
	"log"
	"net"

	"max-db-svc/config"
	"max-db-svc/controller"
	pb "max-db-svc/pb"
	"max-db-svc/repo"
	"max-db-svc/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	svcs  service.Services
	repos repo.Repositories
)

func main() {
	lis, err := net.Listen("tcp", config.Conf.Host)
	log.Print(lis.Addr())
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	initRepos()
	initServices()

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterDatabaseServer(grpcServer, controller.Init(&svcs))
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

func initRepos() {
	repo.InitRepos(&repos)
}

func initServices() {
	svcs.Init(&repos)
}

func init() {
	//load default configuration
	if err := config.LoadConfJson(); err != nil {
		log.Fatalf("Failed loading service config. Error message: %s", err)
	}
	// load environment variables
	config.LoadEnv()
}
