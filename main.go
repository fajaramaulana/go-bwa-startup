package main

import (
	"go-bwa-startup/auth"
	"go-bwa-startup/config"
	"go-bwa-startup/handler"
	"go-bwa-startup/user"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db := config.ConnectionDb()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	modeGin := os.Getenv("GIN_MODE")
	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)
	gin.SetMode(modeGin)
	router := gin.Default()

	api := router.Group("/api/v1/user/")

	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.Login)
	api.POST("/emailchecker", userHandler.CheckEmailAvailability)
	api.POST("/uploadavatar", userHandler.UploadAvatar)

	router.Run(":8081")

	// input dari user
	// handler mapping input dari user menjadi sebuah struct input
	// service mapping dari struct input menjadi struct user
	// repository
}
