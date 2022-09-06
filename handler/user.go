package handler

import (
	"go-bwa-startup/helper"
	"go-bwa-startup/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(uService user.Service) *userHandler {
	return &userHandler{uService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var errors []string
		for _, v := range err.(validator.ValidationErrors) {
			errors = append(errors, v.Error())
		}
		errorMessage := gin.H{"errors": errors}
		jsonResponse := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	// struct diatas kita passing sebagai parameter service Register User
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		var errors []string
		for _, v := range err.(validator.ValidationErrors) {
			errors = append(errors, v.Error())
		}
		errorMessage := gin.H{"errors": errors}
		jsonResponse := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	formatter := user.FormatUserRegister(newUser, "token")

	jsonResponse := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, jsonResponse)

}
