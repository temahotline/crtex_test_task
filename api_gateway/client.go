package api_gateway

import (
	"context"
	"time"

	"api_gateway/protos"
)

func CallGetUser(client protos.UserServiceClient, req *protos.GetUserRequest) (*protos.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CallCreateUser(client protos.UserServiceClient, req *protos.CreateUserRequest) (*protos.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CallGetTransactions(client protos.TransactionServiceClient, req *protos.GetTransactionsRequest) (*protos.Transactions, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.GetTransactions(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CallCreateTransaction(client protos.TransactionServiceClient, req *protos.CreateTransactionRequest) (*protos.Transaction, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.CreateTransaction(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
