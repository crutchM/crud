package http

import "github.com/gin-gonic/gin"

func InitRoutes() *gin.Engine {
	router := gin.New()

	router.POST("/register")

	router.POST("/auth")

	router.GET("/check")

	return router
}
