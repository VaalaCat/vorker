package auth

import (
	"voker/authz"
	"voker/common"
	"voker/defs"
	"voker/models"
	"voker/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoginEndpoint(c *gin.Context) {
	req, err := parseLoginReq(c)
	if err != nil {
		common.RespErr(c, common.RespCodeInvalidRequest,
			common.RespMsgInvalidRequest, nil)
		return
	}

	ok, err := models.CheckUserPassword(req.UserName, req.Password)
	if err != nil || !ok {
		common.RespErr(c, common.RespCodeAuthErr,
			common.RespMsgAuthErr, nil)
		return
	}

	user, err := models.GetUserByUserName(req.UserName)
	if err != nil {
		logrus.WithError(err).Error("get user by user name failed")
		common.RespErr(c, common.RespCodeInternalError,
			common.RespMsgInternalError, nil)
		return
	}

	token, err := utils.SignToken(user.ID)
	if err != nil {
		logrus.WithError(err).Error("sign token failed")
		common.RespErr(c, common.RespCodeInternalError,
			common.RespMsgInternalError, nil)
		return
	}

	authz.SetToken(c, token)

	c.Header(common.AuthorizationHeaderKey, token)
	common.RespOK(c, common.RespMsgOK, defs.LoginResponse{
		Status: common.RespCodeOK,
		Token:  token})
}
