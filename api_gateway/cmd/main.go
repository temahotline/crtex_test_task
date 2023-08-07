package main

import (
	"api_gateway"
	"api_gateway/pkg/handler"
	"api_gateway/protos"
	"google.golang.org/grpc"
	"log"
)

// @title Crtex Test Task
// @version 1.0
// @description API Server for Transactions

// @host localhost:8000
// @BasePath /

func main() {
	conn, err := grpc.Dial("gateway_processor:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect to grpc server: ", err)
	}
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("Failed to close connection: %s", err)
		}
	}()
	userServiceClient := protos.NewUserServiceClient(conn)
	transactionServiceClient := protos.NewTransactionServiceClient(conn)
	handlers := handler.NewHandler(userServiceClient, transactionServiceClient)

	srv := new(api_gateway.Server)
	if err := srv.Start("8000", handlers.InitRouter()); err != nil {
		log.Fatal("failed to start server: ", err)
	}

}
