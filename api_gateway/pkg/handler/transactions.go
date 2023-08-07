package handler

import (
	"api_gateway"
	"api_gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *Handler) getTransactions(c *gin.Context) {
	log.Println("getTransactions called")
	protoReq, err := utils.ParseTransactionsRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := api_gateway.CallGetTransactions(h.transactionServiceClient, protoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) createTransaction(c *gin.Context) {
	log.Println("createTransaction called")
	protoReq, err := utils.ParseTransactionRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := api_gateway.CallCreateTransaction(h.transactionServiceClient, protoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, resp)
}
