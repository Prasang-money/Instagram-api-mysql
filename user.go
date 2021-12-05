package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
}

type Post struct {
	Article string `json:"article,omitempty" bson:"article,omitempty"`
	User    User   `json:"User,omitempty" bson:"User,omitempty"`
}
