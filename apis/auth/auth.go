package auth

import (
	"myauth/service"

	"github.com/gin-gonic/gin"
)

func ApplyRouters(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", service.RegisterUser)
		auth.GET("/login", service.RegisterUser)
	}
}
