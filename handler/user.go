package handler

import (
	"go-bwa-startup/auth"
	"go-bwa-startup/helper"
	"go-bwa-startup/user"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(uService user.Service, aService auth.Service) *userHandler {
	return &userHandler{uService, aService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// tangkap input dari user
	// map input dari user ke struct RegisterUserInput
	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}
		jsonResponse := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	// struct diatas kita passing sebagai parameter service Register User
	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		errors := helper.FormatError(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}
		jsonResponse := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID, newUser.Role)

	if err != nil {
		errors := helper.FormatError(err.(validator.ValidationErrors))
		errorMessage := gin.H{"errors": errors}
		jsonResponse := helper.APIResponse("Register Account Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	formatter := user.FormatUserRegister(newUser, token)

	jsonResponse := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, jsonResponse)

}

func (h *userHandler) Login(c *gin.Context) {
	// user memasukkan input (email & password)
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct passing service
	// di service mencari dg bantuan repository (kalo mvc php = model) user dengan email x
	// mencocokkan password

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		jsonResponse := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	userLogin, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		jsonResponse := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	token, err := h.authService.GenerateToken(userLogin.ID, userLogin.Role)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		jsonResponse := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	jsonResponse := user.FormatUserLogin(userLogin, token)

	c.JSON(http.StatusOK, jsonResponse)

}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// ada input email dari user
	// input email dimapping ke struct input
	// struct input di passing ke service
	// service akan manggil repository email sudah ada atau belum
	// repository - db

	var input user.CheckEmailAvailability

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}

		jsonResponse := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusBadRequest, jsonResponse)
		return
	}

	emailAvailable, err := h.userService.IsEmailAvailable(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		response := helper.APIResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"isAvailable": emailAvailable,
	}

	var metaMessage string

	if emailAvailable {
		metaMessage = "Email is available"
	} else {
		metaMessage = "Email has been registered"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user
	// simpan gambar di folder "images/"
	// di service kita panggil repo
	// JWT (sementara hardcode, seakan akan user yg login ID = 1)
	// repo ambil data user yang ID = 1
	// repo update date user simpan lokasi file di field avatar

	file, err := c.FormFile("avatar")

	if err != nil {
		errors := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload avatar Failed", http.StatusBadRequest, "error", errors)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	extension := filepath.Ext(file.Filename)

	currentTime := file.Filename + time.Now().Format("2006-01-02 15:04:05")

	newFileName := slug.Make(currentTime) + extension

	path := "images/" + newFileName

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		errors := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload avatar Failed", http.StatusBadRequest, "error", errors)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// nanti ini diganti data dari jwt
	userId := 1

	_, err = h.userService.UploadAvatar(userId, path)

	if err != nil {
		errors := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Upload avatar Failed", http.StatusBadRequest, "error", errors)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Upload avatar success", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
	return

}
