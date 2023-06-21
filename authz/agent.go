package authz

import (
	"fmt"
	"vorker/conf"
	"vorker/defs"
	"vorker/utils"

	"github.com/gin-gonic/gin"
)

func AgentAuthz() func(c *gin.Context) {
	return func(c *gin.Context) {
		secret := c.Request.Header.Get(defs.HeaderNodeSecret)
		name := c.Request.Header.Get(defs.HeaderNodeName)
		if secret == "" || name == "" {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		if utils.CheckPasswordHash(
			fmt.Sprintf("%s%s", name, conf.AppConfigInstance.AgentSecret),
			secret) {
			c.Set(defs.KeyNodeName, name)
			c.Next()
			return
		}

		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
	}
}
