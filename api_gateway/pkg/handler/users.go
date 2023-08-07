package handler

import (
	"api_gateway"
	"api_gateway/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TempUserRequest struct {
	UserId int32 `uri:"user_id"`
}

func (h *Handler) signUp(c *gin.Context) {
	log.Println("signUp called")
	req, err := utils.ParseCreateUserRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("signUp ParseCreateUserRequest: %v", err)
		return
	}

	resp, err := api_gateway.CallCreateUser(h.userServiceClient, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		log.Printf("signUp api_gateway.CallCreateUser: %v", err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) getUser(c *gin.Context) {
	log.Println("getUser called")

	protoReq, err := utils.ParseGetUserRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	log.Printf("getUser protoReq: %v", protoReq.UserId)

	resp, err := api_gateway.CallGetUser(h.userServiceClient, protoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
