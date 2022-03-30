package main

import (
	controller "Final-Project-JCC-Golang-2022/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//1. User Biasa 2. Memiliki Toko 3. Admin
	router.POST("/login", controller.UserLogin)
	router.GET("/logout", controller.Logout)

	admin := router.Group("/")
	admin.Use(controller.Authenticate(3))
	{
		router.DELETE("/users/:id", controller.DeleteUser)
	}

	basicUser := router.Group("/")
	basicUser.Use(controller.Authenticate(1))
	{

		router.PUT("/users", controller.UpdateUsers)
	}

	//User
	router.GET("/users", controller.GetAllUsers)
	router.POST("/users", controller.InsertUser)

	router.Run(":8080")

}
