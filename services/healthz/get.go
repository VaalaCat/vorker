package healthz

import "github.com/gin-gonic/gin"

func GetEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
