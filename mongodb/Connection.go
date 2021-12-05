package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb+srv://Prasang:money%409211@cluster0.kts3l.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err1 := client.Ping(ctx, readpref.Primary())

	if err1 != nil {
		log.Fatal("There is a problem  in connection")
	}

	log.Printf("successfully connected to database")

	return client
}
