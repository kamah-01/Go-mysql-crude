package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func main() {
	uri := "mongodb://localhost:27017"
	appCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	client, connectErr := mongo.Connect(appCtx, options.Client().ApplyURI(uri))
	if connectErr != nil {
		panic(connectErr)
	}

	pingErr := client.Ping(appCtx, readpref.Primary())
	if pingErr != nil {
		panic(pingErr)
	}
	fmt.Println("Application connected to mongodb successfully!")

	defer cancel()

	defer func() {
		disconnectErr := client.Disconnect(appCtx)
		if disconnectErr != nil {
			panic(disconnectErr)
		}
	}()
}
