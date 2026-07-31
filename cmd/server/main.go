package main

import (
	"log"
	"net"

	"grpc/internal/database"
	"grpc/internal/handler"
	"grpc/internal/repository"
	"grpc/internal/service"
	pb "grpc/proto"

	"google.golang.org/grpc"
)

func main() {

	// Connect Database
	db := database.Connect()

	// Create Repository
	repo := repository.NewUserRepository(db)

	// Create Service
	userService := service.NewUserService(repo)

	// Create Handler
	userHandler := handler.NewUserHandler(userService)

	// Create Listener
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC Server
	grpcServer := grpc.NewServer()

	// Register Handler
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	log.Println("gRPC Server is running on port 50051")

	// Start Server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}