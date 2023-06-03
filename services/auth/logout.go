package auth

import (
	"voker/common"
	"voker/conf"

	"github.com/gin-gonic/gin"
)

func LogoutEndpoint(c *gin.Context) {
	c.SetCookie(common.AuthorizationKey, "", 0, "",
		conf.AppConfigInstance.CookieDomain, false, true)
	common.RespOK(c, common.RespMsgOK, nil)
}
