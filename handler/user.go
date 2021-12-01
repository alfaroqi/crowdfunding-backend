package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// capture input from user
	// map input from user to struct RegisterUserInput
	// struct above we pass it as a service parameter

	var input user.RegisterUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValitationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Register account failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)

	if err != nil {
		response := helper.ApiResponse("Register account failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// token, err := h.jwtService.GenerateToken(newUser)

	formatter := user.FormatUser(newUser, "tokentokentoken")

	response := helper.ApiResponse("Account has been registered", http.StatusOK, "sucess", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {

	/* 	workflow :

	1. user memasukan input(email, password)
	2. input ditangkap handler
	3. mapping dari input user ke input struct
	5. input struct passing service
	6. di service mencari dengan bantuan repository user dan email x
	7. mencocokan password
	*/

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValitationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedInUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorsMessage := gin.H{"errors": err.Error()}

		response := helper.ApiResponse("Login failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	formatter := user.FormatUser(loggedInUser, "tokentokentoken")

	response := helper.ApiResponse("Logged successfully", http.StatusOK, "sucess", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvaibility(c *gin.Context) {
	/*	workflow :
		- ada input email dari user
		- input email di mapping ke struct input
		- service akan manggil repository - email sudah ada atau belum
		- repository - db
	*/
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValitationError(err)
		errorsMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorsMessage := gin.H{"errors": "Server Error"}

		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}

	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.ApiResponse(metaMessage, http.StatusOK, "sucess", data)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	/*	workflow :
		- input dari user
		- simpan gambarnya ke folder "images/"
		- di service kita panggil repo
		- JWT (sementara handcode, sekaan2 user yang login ID = 1)
		- repo ambil data user dengan ID = 1
		- repo update data user, simpan lokasi file gambar
	*/

	file, err := c.FormFile("avatar")
	if err != nil {
		response := helper.ApiResponse("Upload avatar image failed", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return

	}

	// test pake id manual dulu, harusnya make JWT
	userID := 1

	// images/namafile.png > images/1-namafile.png "1 adalah id user"

	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Upload avatar image failed", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Upload avatar image failed", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return

	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Upload avatar successfully", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
