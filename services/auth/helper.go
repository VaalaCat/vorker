package auth

import (
	"errors"
	"vorker/defs"

	"github.com/gin-gonic/gin"
)

const (
	ErrInvalidRequest = "invalid request"
)

func parseRegisterReq(c *gin.Context) (registerRequest defs.RegisterRequest, err error) {
	registerRequest = defs.RegisterRequest{}
	if err = c.ShouldBindJSON(&registerRequest); err != nil {
		return
	}
	if !registerRequest.Validate() {
		err = errors.New(ErrInvalidRequest)
		return
	}
	return registerRequest, nil
}

func parseLoginReq(c *gin.Context) (loginRequest defs.LoginRequest, err error) {
	loginRequest = defs.LoginRequest{}
	if err = c.ShouldBindJSON(&loginRequest); err != nil {
		return
	}
	if !loginRequest.Validate() {
		err = errors.New(ErrInvalidRequest)
		return
	}
	return loginRequest, nil
}
