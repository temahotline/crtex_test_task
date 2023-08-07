package handler

import (
	"api_gateway"
	"api_gateway/protos"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TempUserRequest struct {
	UserId int32 `uri:"user_id"`
}

func (h *Handler) signUp(c *gin.Context) {
	log.Println("signUp called")
	var req protos.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("signUp c.BindJSON: %v", err)
		return
	}

	resp, err := api_gateway.CallCreateUser(h.userServiceClient, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		log.Printf("signUp api_gateway.CallCreateUser: %v", err)
		return
	}

	c.JSON(http.StatusCreated, resp)
}

func (h *Handler) getUser(c *gin.Context) {
	log.Println("getUser called")
	//var req protos.GetUserRequest
	var tempReq TempUserRequest

	if err := c.BindUri(&tempReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	log.Printf("getUser req: %v", tempReq.UserId)

	var protoReq protos.GetUserRequest
	protoReq.UserId = tempReq.UserId
	log.Printf("getUser protoReq: %v", protoReq.UserId)

	resp, err := api_gateway.CallGetUser(h.userServiceClient, &protoReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}
