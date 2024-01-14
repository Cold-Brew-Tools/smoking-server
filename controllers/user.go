package controllers

import (
	"cold-brew-smoking-server/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var UsersCollection *mongo.Collection

func AddUserEndpoints(engine *gin.Engine, database *mongo.Database) {
	UsersCollection = database.Collection("Users")
	usersGroup := engine.Group("/users")
	usersGroup.POST("/", registerUser)
}

func registerUser(ginContext *gin.Context) {
	requestBody := &models.RegisterUserBody{}
	ginContext.BindJSON(requestBody)

	userDocument, httpRequestError := requestBody.ToUserDocument()
	if httpRequestError != nil {
		ginContext.JSON(httpRequestError.Status, httpRequestError)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	UsersCollection.InsertOne(ctx, userDocument)

	ginContext.JSON(http.StatusCreated, userDocument)
}
