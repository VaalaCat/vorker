package authz

import (
	"fmt"
	"vorker/conf"
	"vorker/defs"
	sec "vorker/utils/secret"

	"github.com/gin-gonic/gin"
)

func AgentAuthz() func(c *gin.Context) {
	return func(c *gin.Context) {
		secret := c.Request.Header.Get(defs.HeaderNodeSecret)
		name := c.Request.Header.Get(defs.HeaderNodeName)
		querySec := c.Copy().Query("secret")
		queryName := c.Copy().Query("name")
		if (secret == "" || name == "") && (querySec == "" || queryName == "") {
			c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		if sec.CheckPasswordHash(
			fmt.Sprintf("%s%s", name, conf.AppConfigInstance.AgentSecret),
			secret) {
			c.Set(defs.KeyNodeName, name)
			c.Next()
			return
		}

		if sec.CheckPasswordHash(
			fmt.Sprintf("%s%s", queryName, conf.AppConfigInstance.AgentSecret),
			querySec) {
			c.Set(defs.KeyNodeName, queryName)
			c.Next()
			return
		}

		c.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
	}
}
