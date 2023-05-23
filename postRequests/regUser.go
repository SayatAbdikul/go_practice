package postrequests

import (
	"context"
	"log"
	"net/http"

	"github.com/SayatAbdikul/go_practice/server"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}

func RegUser(c *gin.Context) {
	collection := server.Database.Collection("users")
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		message := "the json request is incorrect"
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": message})
	}
	insertResult, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Inserted document ID: ", insertResult.InsertedID)
}
