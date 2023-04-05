package main

import (
	"github.com/gin-gonic/gin"
	"lemonadee/server"
)

func main() {
	r := NewGinRouter()
	r.POST("/users/create", server.CreateUser)
	r.POST("/transactions/create", server.CreateTransaction)
	r.GET("/users", server.GetAllUsers)
	go server.Consume()
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func NewGinRouter() *gin.Engine {
	ginRouter := gin.Default()
	return ginRouter
}
