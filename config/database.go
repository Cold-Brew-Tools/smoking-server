package config

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoOpenConnection(username string, password string, host string, port int) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI("mongodb://"+username+":"+password+"@"+host+":"+strconv.Itoa(port)),
	)

	if err != nil {
		fmt.Println(fmt.Errorf("%w", err))
		panic("Could not open a connection to your database. Check the details.")
	}

	return client
}

func MongoCloseConnection(client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}
