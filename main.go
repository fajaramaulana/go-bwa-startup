package main

import (
	"fmt"
	"go-bwa-startup/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/go-startup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	userInput := user.RegisterUserInput{}

	userInput.Name = "Fajar"
	userInput.Occupation = "Programmer"
	userInput.Email = "fajar@gmail.com"
	userInput.Password = "rahasia"

	newUser, err := userService.RegisterUser(userInput)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%# v", newUser)

	// input dari user
	// handler mapping input dari user menjadi sebuah struct input
	// service mapping dari struct input menjadi struct user
	// repository
}
