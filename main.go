package main

import (
	"bookstore-user-api/app"
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func main(){
	log.Println("start application!")
	app.StartApplication()
}