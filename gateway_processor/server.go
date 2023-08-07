package gateway_processor

import (
	"gateway_processor/pkg/transaction"
	"gateway_processor/pkg/user"
	"gateway_processor/protos"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log"
	"net"
)

func NewUserServer(db *pgxpool.Pool) *user.UserServiceServer {
	return &user.UserServiceServer{DB: db}
}

func NewTransactionServer(db *pgxpool.Pool) *transaction.TransactionServiceServer {
	return &transaction.TransactionServiceServer{DB: db}
}

func StartGRPCServer(address string, db *pgxpool.Pool) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	protos.RegisterUserServiceServer(s, NewUserServer(db))
	protos.RegisterTransactionServiceServer(s, NewTransactionServer(db))

	log.Printf("Starting gRPC server on %s", address)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
