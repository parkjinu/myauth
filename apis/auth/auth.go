package auth

import (
	"myauth/service"

	"github.com/gin-gonic/gin"
)

func ApplyRouters(r *gin.RouterGroup) {
	h := &service.AuthService{}
	auth := r.Group("/auth")
	{
		auth.POST("/register", h.RegisterUser)
		auth.GET("/login", h.RegisterUser)
	}
}
