package main

import (
	api "myauth/apis"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	api.ApplyRouters(router)

	router.Run(":8090")
}
