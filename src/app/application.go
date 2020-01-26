package app

import (
	"bookstore-oauth-api/src/domain/access_token"
	"bookstore-oauth-api/src/http"
	"bookstore-oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	g = gin.Default()
)

func StartApplication() {
	atService := access_token.NewService(db.New())
	atHandler := http.NewHandler(atService)

	g.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	g.POST("/oauth/access_token", atHandler.Create)
	g.PUT("/oauth/access_token/:access_token_id", atHandler.UpdateExpires)
	g.Run(":5002")
}
