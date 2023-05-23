package main

import (
	"fmt"

	post "github.com/SayatAbdikul/go_practice/postRequests"
	"github.com/SayatAbdikul/go_practice/server"
	"github.com/gin-gonic/gin"
)

func main() {
	server.Connect()
	// Connected successfully
	fmt.Println("Connected to MongoDB!")
	router := gin.Default()
	router.POST("/reg_user", post.RegUser)
	router.Run(":6666")
}
