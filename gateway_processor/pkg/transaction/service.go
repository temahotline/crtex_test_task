package transaction

import (
	"context"
	"gateway_processor/pkg/database"
	"gateway_processor/protos"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type TransactionServiceServer struct {
	protos.UnimplementedTransactionServiceServer
	DB *pgxpool.Pool
}

func (s *TransactionServiceServer) GetTransactions(ctx context.Context, req *protos.GetTransactionsRequest) (*protos.Transactions, error) {
	log.Println("GetTransactions called!")
	return database.GetTransactions(ctx, s.DB, req.UserId)
}

func (s *TransactionServiceServer) CreateTransaction(ctx context.Context, req *protos.CreateTransactionRequest) (*protos.Transaction, error) {
	log.Println("CreateTransaction called!")
	return database.CreateTransaction(ctx, s.DB, req.UserId, req.Amount, req.TransactionType)
}
