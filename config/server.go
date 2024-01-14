package config

import (
	"cold-brew-smoking-server/controllers"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func StartGinServer(database *mongo.Database, port int) {
	engine := gin.Default()
	// gin.SetMode(gin.ReleaseMode)

	controllers.AddUserEndpoints(engine, database)

	engine.Run(":" + strconv.Itoa(port))
}
