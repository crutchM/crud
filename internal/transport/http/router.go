package http

import (
	"crud/internal/core/interface/service"
	"crud/internal/transport/handler"
	"github.com/gin-gonic/gin"
)

func InitRoutes(service service.AuthService) *gin.Engine {
	router := gin.New()

	router.POST("/register", handler.RegisterUser(service))

	router.POST("/auth/:id")

	router.GET("/check")

	return router
}
