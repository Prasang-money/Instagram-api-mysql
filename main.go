package main

import (
	"project/Go_db_Api/mongodb"

	"github.com/gin-gonic/gin"
)

var client = mongodb.ConnectDB()
var collection1 = client.Database("go-rest-api").Collection("user")
var collection2 = client.Database("go-rest-api").Collection("Post")

func main() {
	router := gin.Default()

	router.GET("/", health_checkup)
	router.POST("/users", create_user)
	router.GET("/users/:id", GetUserByID)
	router.POST("/posts", UserPost)
	router.Run(":8080")
}
