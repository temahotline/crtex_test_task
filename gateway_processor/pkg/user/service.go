package user

import (
	"context"
	"gateway_processor/pkg/database"
	"gateway_processor/protos"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type UserServiceServer struct {
	protos.UnimplementedUserServiceServer
	DB *pgxpool.Pool
}

func (s *UserServiceServer) CreateUser(ctx context.Context, req *protos.CreateUserRequest) (*protos.User, error) {
	log.Println("CreateUser called!")
	return database.CreateUser(ctx, s.DB, req.FirstName, req.LastName, req.Balance)
}

func (s *UserServiceServer) GetUser(ctx context.Context, req *protos.GetUserRequest) (*protos.User, error) {
	log.Println("GetUser called!")
	return database.GetUser(ctx, s.DB, req.UserId)
}
