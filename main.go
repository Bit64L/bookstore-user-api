package main

import (
	"bookstore-user-api/app"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func main(){
	app.StartApplication()
}