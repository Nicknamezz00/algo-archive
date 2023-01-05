package app

import (
	"algo-archive/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type Response struct {
	Ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

func (r *Response) Reply(data interface{}) {
	hostname, _ := os.Hostname()
	if nil == data {
		data = gin.H{
			"code":  0,
			"msg":   "success",
			"trace": hostname,
		}
	} else {
		data = gin.H{
			"code":  0,
			"msg":   "success",
			"data":  data,
			"trace": hostname,
		}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ReplyError(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
