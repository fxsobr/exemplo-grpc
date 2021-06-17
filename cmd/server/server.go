package main

import (
	"log"
	"net"

	"github.com/fxsobr/exemplo-grpc/pb/pb"
	"github.com/fxsobr/exemplo-grpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	grcpServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grcpServer, services.NewUserService())
	reflection.Register(grcpServer)

	if err := grcpServer.Serve(lis); err != nil {
		log.Fatalf("Could not serve: %¨v", err)
	}

}
