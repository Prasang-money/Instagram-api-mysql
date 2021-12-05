package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func health_checkup(c *gin.Context) {
	c.JSON(http.StatusOK, "Api is Running..........")
}

func create_user(c *gin.Context) {

	var user User
	c.BindJSON(&user)
	result, err := collection1.InsertOne(context.TODO(), user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)

	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": result,
	})

}

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	var user User

	Id, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": Id}
	err := collection1.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UserPost(c *gin.Context) {
	var post Post
	c.BindJSON(&post)
	result, err := collection2.InsertOne(context.TODO(), post)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		log.Fatal(err)

	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": result,
	})

}
