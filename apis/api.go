package api

import (
	"myauth/apis/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRouters(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/health", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "UP",
			})
		})

		auth.ApplyRouters(api)
	}
}
