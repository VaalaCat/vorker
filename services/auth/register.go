package auth

import (
	"voker/common"
	"voker/conf"
	"voker/defs"
	"voker/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterEndpoint(c *gin.Context) {
	if !conf.AppConfigInstance.EnableRegister {
		if count, err := models.AdminGetUserNumber(); err != nil {
			common.RespErr(c, common.RespCodeInternalError,
				common.RespMsgInternalError, nil)
			return
		} else if count >= 1 {
			common.RespErr(c, common.RespCodeMethodNotAllowed,
				common.RespMsgMethodNotAllowed, nil)
			return
		}
	}

	// get userName and email and password from request body and validate them
	req, err := parseRegisterReq(c)
	if err != nil {
		common.RespErr(c, common.RespCodeInvalidRequest,
			common.RespMsgInvalidRequest, nil)
		return
	}

	// check if userName or email already exists
	if err := models.CheckUserNameAndEmail(req.UserName, req.Email); err != nil && err != gorm.ErrRecordNotFound {
		common.RespErr(c, common.RespCodeUserAlreadyExists,
			common.RespMsgUserAlreadyExists, nil)
		return
	}

	// create user
	user := &models.User{
		UserName: req.UserName,
		Password: req.Password,
		Email:    req.Email,
		Status:   common.UserStatusPending,
		Role:     common.UserRoleNormal,
	}
	if err := models.CreateUser(user); err != nil {
		common.RespErr(c, common.RespCodeInternalError,
			common.RespMsgInternalError, nil)
		return
	}

	common.RespOK(c, common.RespMsgOK, &defs.RegisterResponse{
		Status: common.RespCodeOK,
	})
}
