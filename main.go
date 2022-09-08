package main

import (
	"go-bwa-startup/config"
	"go-bwa-startup/handler"
	"go-bwa-startup/user"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectionDb()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/user/login", userHandler.Login)

	router.Run(":8081")

	// input dari user
	// handler mapping input dari user menjadi sebuah struct input
	// service mapping dari struct input menjadi struct user
	// repository
}
