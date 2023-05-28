package main

import (
	"fmt"

	"github.com/SayatAbdikul/go_practice/server"
	user "github.com/SayatAbdikul/go_practice/userRequests"
	"github.com/gin-gonic/gin"
)

func main() {
	server.Connect()
	// Connected successfully
	fmt.Println("Connected to MongoDB!")
	router := gin.Default()
	router.POST("/reg_user", user.RegUser)
	router.PUT("/update_user/:id", user.UpdateUser)
	router.Run(":6666")
}
