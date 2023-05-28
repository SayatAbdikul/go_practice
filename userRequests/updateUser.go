package userRequests

import (
	"context"
	"fmt"
	"net/http"

	"github.com/SayatAbdikul/go_practice/server"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUser(c *gin.Context) {
	collection := server.Database.Collection("users")
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	ctx := context.TODO()
	id, _ := primitive.ObjectIDFromHex(c.Param("id"))
	filter := bson.D{{"_id", id}}
	update := bson.D{{"$set", bson.D{{"name", user.Name}, {"login", user.Login}, {"age", user.Age}, {"password", user.Password}}}}
	fmt.Println(id)
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}
	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
