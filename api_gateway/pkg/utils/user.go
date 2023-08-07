package utils

import (
	"api_gateway/protos"
	"github.com/gin-gonic/gin"
)

type TempUserRequest struct {
	UserId int32 `uri:"user_id"`
}

func ParseCreateUserRequest(c *gin.Context) (*protos.CreateUserRequest, error) {
	var req protos.CreateUserRequest

	if err := c.BindJSON(&req); err != nil {
		return nil, err
	}

	return &req, nil
}

func ParseGetUserRequest(c *gin.Context) (*protos.GetUserRequest, error) {
	var tempReq TempUserRequest

	if err := c.BindUri(&tempReq); err != nil {
		return nil, err
	}

	var protoReq protos.GetUserRequest
	protoReq.UserId = tempReq.UserId

	return &protoReq, nil
}
