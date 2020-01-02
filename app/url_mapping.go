package app

import "bookstore-user-api/controllers"

func mapUrls() {
	router.GET("/ping", controllers.Ping)

	router.GET("/users/:user_id", controllers.GetUser)
	//router.GET("/users/search", controllers.SearchUser)
	router.POST("/user", controllers.CreateUser)

}