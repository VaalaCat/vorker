package auth

import (
	"voker/common"
	"voker/defs"
	"voker/models"

	"github.com/gin-gonic/gin"
)

func GetUserEndpoint(c *gin.Context) {
	uid := c.GetUint(common.UIDKey)
	user, err := models.GetUserByUserID(uid)
	if err != nil {
		common.RespErr(c, common.RespCodeDBErr, common.RespMsgDBErr, nil)
		return
	}
	common.RespOK(c, "ok", &defs.GetUserResponse{
		UserName: user.UserName,
		Role:     user.Role,
		Email:    user.Email,
	})
}
