package auth

import (
	"errors"
	"vorker/entities"

	"github.com/gin-gonic/gin"
)

const (
	ErrInvalidRequest = "invalid request"
)

func parseRegisterReq(c *gin.Context) (registerRequest entities.RegisterRequest, err error) {
	registerRequest = entities.RegisterRequest{}
	if err = c.ShouldBindJSON(&registerRequest); err != nil {
		return
	}
	if !registerRequest.Validate() {
		err = errors.New(ErrInvalidRequest)
		return
	}
	return registerRequest, nil
}

func parseLoginReq(c *gin.Context) (loginRequest entities.LoginRequest, err error) {
	loginRequest = entities.LoginRequest{}
	if err = c.ShouldBindJSON(&loginRequest); err != nil {
		return
	}
	if !loginRequest.Validate() {
		err = errors.New(ErrInvalidRequest)
		return
	}
	return loginRequest, nil
}
