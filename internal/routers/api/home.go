package api

import (
	"algo-archive/pkg/app"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Version(c *gin.Context) {
	response := app.NewResponse(c)
	response.Ctx.JSON(http.StatusOK, gin.H{
		"Version": "v1.0",
	})
}
