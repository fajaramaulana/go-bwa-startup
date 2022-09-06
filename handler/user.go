package handler

import (
	"go-bwa-startup/helper"
	"go-bwa-startup/user"
	"net/http"

	"github.com/gin-gonic/gin"
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
		jsonResponse := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, jsonResponse)
	}

	// struct diatas kita passing sebagai parameter service Register User
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		jsonResponse := helper.APIResponse(err.Error(), http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, jsonResponse)
	}

	formatter := user.FormatUserRegister(newUser, "token")

	jsonResponse := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, jsonResponse)

}
