package authz

import (
	"strings"
	"time"
	"voker/common"
	"voker/conf"
	"voker/utils"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware check if authed and set uid to context
func JWTMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		cookieToken, err := c.Cookie(conf.AppConfigInstance.CookieName)
		if err == nil {
			if t, err := utils.ParseToken(cookieToken); err == nil {
				c.Set(common.UIDKey, t.UID)
				resignJWT(c, t)
				c.Next()
				return
			}
		}

		tokenOrigin := c.Request.Header.Get(common.AuthorizationKey)
		tokenList := strings.Split(tokenOrigin, " ")
		if len(tokenList) != 2 {
			common.RespErr(c, common.RespCodeNotAuthed, common.RespMsgNotAuthed, nil)
			c.Abort()
			return
		}
		tokenStr := tokenList[1]

		if tokenStr == "" {
			common.RespErr(c, common.RespCodeNotAuthed, common.RespMsgNotAuthed, nil)
			c.Abort()
			return
		}

		if t, err := utils.ParseToken(tokenStr); err == nil {
			c.Set(common.UIDKey, t.UID)
			resignJWT(c, t)
			c.Next()
			return
		}

		common.RespErr(c, common.RespCodeNotAuthed, common.RespMsgNotAuthed, nil)
		c.Abort()
	}
}

func resignJWT(c *gin.Context, t *conf.JwtClaims) error {
	if time.Until(t.ExpiresAt.Time) > time.Duration(conf.AppConfigInstance.CookieAge/2) {
		return nil
	}

	token, err := utils.SignToken(t.UID)
	if err != nil {
		return err
	}
	c.SetCookie(conf.AppConfigInstance.CookieName,
		token,
		conf.AppConfigInstance.CookieAge,
		"/",
		conf.AppConfigInstance.CookieDomain,
		true, true)
	return nil
}

func SetToken(c *gin.Context, token string) {
	c.SetCookie(conf.AppConfigInstance.CookieName,
		token,
		conf.AppConfigInstance.CookieAge,
		"/",
		conf.AppConfigInstance.CookieDomain,
		true, true)
}
