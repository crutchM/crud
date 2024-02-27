package handler

import (
	"crud/internal/core/interface/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userhttp struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func RegisterUser(service service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user userhttp

		if err := c.BindJSON(&user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": "неверное тело запроса"})

			return
		}

		token, err := service.Register(c.Request.Context(), user.Login, user.Password)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{"message": err.Error()})

			return
		}

		c.JSON(http.StatusOK, token)
	}
}
