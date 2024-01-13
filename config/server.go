package config

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func StartGinServer(port int) {
	engine := gin.Default()
	// gin.SetMode(gin.ReleaseMode)

	engine.Run(":" + strconv.Itoa(port))
}
