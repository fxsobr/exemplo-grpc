package main

import (
	"context"
	"log"

	"github.com/fxsobr/exemplo-grpc/pb/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC server: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
}

func AddUser(client pb.UserServiceClient) {

	req := &pb.User{
		Id:    "0",
		Name:  "Marcelo Ceolin",
		Email: "mceolin@unidavi.edu.br",
	}

	res, err := client.AddUser(context.Background(), req)

}
