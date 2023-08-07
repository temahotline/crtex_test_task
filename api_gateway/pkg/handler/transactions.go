package handler

import (
	"api_gateway"
	"api_gateway/protos"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TempTransactionsRequest struct {
	UserId int32 `uri:"user_id" binding:"required"`
}

type TempCreateTransactionRequestJSON struct {
	Amount          int32  `json:"amount" binding:"required"`
	TransactionType string `json:"transaction_type" binding:"required"`
}

type TempCreateTransactionRequest struct {
	UserId int32 `uri:"user_id" binding:"required"`
}

func (h *Handler) getTransactions(c *gin.Context) {
	log.Println("getTransactions called")
	var tempReq TempTransactionsRequest

	if err := c.BindUri(&tempReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var protoReq protos.GetTransactionsRequest
	protoReq.UserId = tempReq.UserId

	resp, err := api_gateway.CallGetTransactions(h.transactionServiceClient, &protoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) createTransaction(c *gin.Context) {
	log.Println("createTransaction called")
	var tempJson TempCreateTransactionRequestJSON
	var tempReq TempCreateTransactionRequest

	if err := c.BindUri(&tempReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("tempReq: %+v\n", tempReq)

	if err := c.BindJSON(&tempJson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("tempReq2: %+v\n", tempReq)

	var protoReq protos.CreateTransactionRequest
	protoReq.UserId = tempReq.UserId
	protoReq.Amount = tempJson.Amount
	protoReq.TransactionType = tempJson.TransactionType

	resp, err := api_gateway.CallCreateTransaction(h.transactionServiceClient, &protoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}
