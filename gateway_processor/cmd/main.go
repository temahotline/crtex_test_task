package main

import (
	"gateway_processor"
	"gateway_processor/pkg/database"
	"log"
)

const (
	port = ":50051"
)

func main() {
	log.Println("Starting application...")

	connString := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
	pool, err := database.NewConnection(connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()

	gateway_processor.StartGRPCServer(port, pool)
}
