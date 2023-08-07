package database

import (
	"context"
	"fmt"
	"gateway_processor/protos"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

func CreateUser(
	ctx context.Context, db *pgxpool.Pool, firstName string, lastName string, balance int32) (*protos.User, error) {

	query := "INSERT INTO users (first_name, last_name, balance) VALUES ($1, $2, $3) RETURNING id"
	var id int32
	err := db.QueryRow(ctx, query, firstName, lastName, balance).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &protos.User{Id: id, FirstName: firstName, LastName: lastName, Balance: balance}, nil
}

func GetUser(ctx context.Context, db *pgxpool.Pool, userId int32) (*protos.User, error) {
	log.Printf("GetUser called with ID %d", userId)
	query := "SELECT first_name, last_name, balance FROM users WHERE id = $1"
	var firstName string
	var lastName string
	var balance int32
	err := db.QueryRow(ctx, query, userId).Scan(&firstName, &lastName, &balance)
	if err != nil {
		log.Printf("Error querying user with ID %d: %v", userId, err)
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user with ID %d not found", userId)
		}
		return nil, fmt.Errorf("database error: %v", err)
	}
	return &protos.User{Id: userId, FirstName: firstName, LastName: lastName, Balance: balance}, nil
}
