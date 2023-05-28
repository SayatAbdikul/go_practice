package userRequests

import (
	"context"
	"log"
	"net/http"

	"github.com/SayatAbdikul/go_practice/server"
	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string `json:"name" bson:"name"`
	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`
	Age      int    `json:"age" bson:"age"`
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
