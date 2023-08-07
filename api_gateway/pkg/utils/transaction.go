package utils

import (
	"api_gateway/protos"
	"github.com/gin-gonic/gin"
)

type TempCreateTransactionRequestJSON struct {
	Amount          int32  `json:"amount" binding:"required"`
	TransactionType string `json:"transaction_type" binding:"required"`
}

type TempCreateTransactionRequest struct {
	UserId int32 `uri:"user_id" binding:"required"`
}

type TempTransactionsRequest struct {
	UserId int32 `uri:"user_id" binding:"required"`
}

func ParseTransactionRequest(c *gin.Context) (*protos.CreateTransactionRequest, error) {
	var tempJson TempCreateTransactionRequestJSON
	var tempReq TempCreateTransactionRequest

	if err := c.BindUri(&tempReq); err != nil {
		return nil, err
	}

	if err := c.BindJSON(&tempJson); err != nil {
		return nil, err
	}

	return &protos.CreateTransactionRequest{
		UserId:          tempReq.UserId,
		Amount:          tempJson.Amount,
		TransactionType: tempJson.TransactionType,
	}, nil
}

func ParseTransactionsRequest(c *gin.Context) (*protos.GetTransactionsRequest, error) {
	var tempReq TempTransactionsRequest

	if err := c.BindUri(&tempReq); err != nil {
		return nil, err
	}

	return &protos.GetTransactionsRequest{
		UserId: tempReq.UserId,
	}, nil
}
