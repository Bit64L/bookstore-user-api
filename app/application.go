package app

import (
	"bookstore-user-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	logger.Info("start the application")
	mapUrls()
	router.Run(":5001")
}
