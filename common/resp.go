package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

func (r *Response) WithCode(code int) *Response {
	r.Code = code
	return r
}

func (r *Response) WithMsg(msg string) *Response {
	r.Msg = msg
	return r
}

func (r *Response) WithData(data interface{}) *Response {
	r.Data = data
	return r
}

func RespOK(c *gin.Context, msg string, data interface{}) {
	logrus.WithContext(c).Infof(msg)
	c.JSON(http.StatusOK, &Response{Code: RespCodeOK, Msg: msg, Data: data})
}

func RespErr(c *gin.Context, code int, errMsg string, data interface{}) {
	logrus.WithContext(c).Errorf(errMsg)
	c.JSON(http.StatusOK, &Response{Code: code, Msg: errMsg, Data: data})
}
