package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ramnakm/account-service/database"
	"github.com/ramnakm/account-service/handlers"
)

func main() {
	database.Connect()
	r := gin.Default()
	r.POST("/accounts", handlers.CreateAccount)
	r.GET("/accounts/:id", handlers.GetAccount)
	r.Run(":8080")
}
