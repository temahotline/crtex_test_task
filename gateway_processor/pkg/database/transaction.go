package database

import (
	"context"
	"errors"
	"gateway_processor/protos"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func CreateTransaction(
	ctx context.Context,
	db *pgxpool.Pool,
	userId int32,
	amount int32,
	transactionType string) (*protos.Transaction, error) {

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, "SET TRANSACTION ISOLATION LEVEL SERIALIZABLE")

	defer func() {
		if rErr := tx.Rollback(ctx); rErr != nil && rErr != pgx.ErrTxClosed {
			log.Printf("Failed to rollback: %v", rErr)
		}
	}()

	// Получить текущий баланс пользователя
	var currentBalance int32
	err = tx.QueryRow(ctx, "SELECT balance FROM users WHERE id = $1", userId).Scan(&currentBalance)
	if err != nil {
		return nil, err
	}

	var balanceBefore int32
	switch transactionType {
	case "deposit":
		balanceBefore = currentBalance
		currentBalance += amount
	case "withdrawal":
		balanceBefore = currentBalance
		currentBalance -= amount
		if currentBalance < 0 {
			return nil, errors.New("insufficient funds")
		}
	default:
		return nil, errors.New("unknown transaction type")
	}

	_, err = tx.Exec(ctx, "UPDATE users SET balance = $1 WHERE id = $2", currentBalance, userId)
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO transactions (user_id, amount, balance_before, transaction_type) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int32
	err = tx.QueryRow(ctx, query, userId, amount, balanceBefore, transactionType).Scan(&id)
	if err != nil {
		return nil, err
	}

	if err = tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &protos.Transaction{
		Id:              id,
		UserId:          userId,
		Amount:          amount,
		BalanceBefore:   balanceBefore,
		TransactionType: transactionType,
	}, nil
}

func GetTransactions(ctx context.Context, db *pgxpool.Pool, userId int32) (*protos.Transactions, error) {
	query := "SELECT id, amount, balance_before, transaction_type FROM transactions WHERE user_id = $1"
	rows, err := db.Query(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*protos.Transaction
	for rows.Next() {
		var id int32
		var amount int32
		var balanceBefore int32
		var transactionType string
		err := rows.Scan(&id, &amount, &balanceBefore, &transactionType)
		if err != nil {
			return nil, err
		}
		transactions = append(
			transactions,
			&protos.Transaction{
				Id:              id,
				UserId:          userId,
				Amount:          amount,
				BalanceBefore:   balanceBefore,
				TransactionType: transactionType,
			})
	}

	return &protos.Transactions{Transactions: transactions}, nil
}
