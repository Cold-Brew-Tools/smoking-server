package main

import (
	"cold-brew-smoking-server/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var SmokingDatabase *mongo.Database

func main() {
	mongoClient := config.MongoOpenConnection("root", "example", "localhost", 27017)
	SmokingDatabase = mongoClient.Database("SmokingDatabase")
	defer config.MongoCloseConnection(mongoClient)

	config.StartGinServer(8080)
}
