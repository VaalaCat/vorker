package request

import (
	"errors"
	"voker/common"

	"github.com/gin-gonic/gin"
)

func Bind[T common.Request](c *gin.Context, req T) (err error) {
	if err := c.ShouldBindJSON(req); err != nil {
		return err
	}
	if !req.Validate() {
		return errors.New(common.ErrMsgInvalidRequest)
	}
	return nil
}
