package auth

import (
	"runtime/debug"
	"vorker/common"
	"vorker/entities"
	"vorker/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetUserEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	uid := c.GetUint(common.UIDKey)
	user, err := models.GetUserByUserID(uid)
	if err != nil {
		common.RespErr(c, common.RespCodeDBErr, common.RespMsgDBErr, nil)
		return
	}
	common.RespOK(c, "ok", &entities.GetUserResponse{
		UserName: user.UserName,
		Role:     user.Role,
		Email:    user.Email,
	})
}
