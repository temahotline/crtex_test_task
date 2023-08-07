package handler

import (
	"api_gateway/protos"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	userServiceClient        protos.UserServiceClient
	transactionServiceClient protos.TransactionServiceClient
}

func (h *Handler) InitRouter() *gin.Engine {
	route := gin.New()

	api := route.Group("/api")
	{
		users := api.Group("/users")
		{
			users.POST("/", h.signUp)
			users.GET("/:user_id", h.getUser)
			transactions := users.Group("/:user_id/transactions")
			{
				transactions.POST("/", h.createTransaction)
				transactions.GET("/", h.getTransactions)
			}
		}
	}
	return route
}

func NewHandler(
	userServiceClient protos.UserServiceClient,
	transactionServiceClient protos.TransactionServiceClient) *Handler {
	return &Handler{
		userServiceClient:        userServiceClient,
		transactionServiceClient: transactionServiceClient,
	}
}
