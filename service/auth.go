package service

import (
	"github.com/gin-gonic/gin"
)

func RegisterUser(ctx *gin.Context) {
	// TODO: insert daba to database
	ctx.JSON(200, "success")
}
