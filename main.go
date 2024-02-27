package main

import (
	api "myauth/apis"
	db "myauth/config"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectMySQL()
}

func main() {

	router := gin.Default()

	api.ApplyRouters(router)

	router.Run(":8090")
}
