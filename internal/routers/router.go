package routers

import (
	"algo-archive/internal/routers/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewRouter() *gin.Engine {
	e := gin.New()
	e.HandleMethodNotAllowed = true
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	// Cross region conf
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")
	e.Use(cors.New(corsConfig))

	// v1 group api
	r := e.Group("/v1")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK v1",
		})
	})

	// Register
	r.POST("/auth/register", api.Register)

	// Login
	r.POST("/auth/login", api.Login)

	// no auth api
	//noAuthAPI := r.Group("/")
	//{
	//	noAuthAPI.GET("/", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"message": "OK",
	//		})
	//	})
	//}
	return e
}
