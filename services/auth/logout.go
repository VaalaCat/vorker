package auth

import (
	"runtime/debug"
	"vorker/common"
	"vorker/conf"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogoutEndpoint(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("Recovered in f: %+v, stack: %+v", r, string(debug.Stack()))
			common.RespErr(c, common.RespCodeInternalError, common.RespMsgInternalError, nil)
		}
	}()
	c.SetCookie(common.AuthorizationKey, "", 0, "",
		conf.AppConfigInstance.CookieDomain, false, true)
	common.RespOK(c, common.RespMsgOK, nil)
}
