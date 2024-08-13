package handler

import (
	"net/http"
	"tesptzen/cmd/appuser"
	"tesptzen/cmd/auth"
	"tesptzen/helpers"
	"tesptzen/utilitys"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	authService auth.IAuthServices
}

func NewUserHandler(authService auth.IAuthServices) *UserHandler {
	return &UserHandler{authService}
}

type H map[string]interface{}

func (u *UserHandler) RegisterUser(c *gin.Context) {
	var input appuser.RegisterUserInputDto
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Register Failed", http.StatusUnprocessableEntity, "error", errormessage)
		c.JSON(http.StatusUnprocessableEntity, response)

	}

	applicationusersschan := make(chan appuser.ApplicationUser)
	errorChan := make(chan error)
	go func() {
		appss, errchan := u.authService.RegisterUser(input)
		if errchan != nil {
			errorChan <- errchan
			return
		}
		applicationusersschan <- appss
	}()

	var userx appuser.ApplicationUser
	var errh error

	// Menunggu hasil dari kedua goroutine
	for i := 0; i < 1; i++ {
		select {
		case userx = <-applicationusersschan:
		case errh = <-errorChan:
		}
	}

	//_, err2 := u.authService.RegisterUser(input, &wg)
	if errh != nil {
		utilitys.LogError(errh)
		errormessage := H{"errors": errh.Error()}
		response := helpers.BaseCommandResponse("Register Failed", http.StatusBadRequest, "error", errormessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helpers.BaseCommandResponse("Account has been registered", http.StatusOK, "success", userx)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) Login(c *gin.Context) {
	var input appuser.LoginDto
	err := c.ShouldBindJSON(&input)
	if err != nil {
		utilitys.LogError(err)
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Failed", http.StatusUnprocessableEntity, "error", errormessage)

		c.JSON(http.StatusUnprocessableEntity, response)
	}
	// appuserIsExist, _ := u.userServices.FindApplicationuserByName(input.Username)
	// if appuserIsExist.Id == 0 {
	// 	errorMessage := H{"errors": "User Not Found"}
	// 	response := helpers.BaseCommandResponse("User Not Found", http.StatusBadRequest, "error", errorMessage)

	// 	c.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	token, err2 := u.authService.LoginCheck(input.Username, input.Password)
	if err2 != nil {
		utilitys.LogError(err2)
		errorMessage := H{"errors": token}
		response := helpers.BaseCommandResponse("User Not Found", http.StatusBadRequest, "error", errorMessage)

		c.JSON(http.StatusBadRequest, response)
	}
	response := helpers.BaseCommandResponse("success", http.StatusOK, "success", token)
	c.JSON(http.StatusOK, response)
}

func (u *UserHandler) ResetPasswordUser(c *gin.Context) {
	var input appuser.ResetPasswordDto
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helpers.FormatError(err)
		errormessage := H{"errors": errors}
		response := helpers.BaseCommandResponse("Failed", http.StatusUnprocessableEntity, "error", errormessage)

		c.JSON(http.StatusUnprocessableEntity, response)
	}
	asd, err2 := u.authService.ResetPassword(input)
	if err2 != nil {
		errormessage := H{"errors": err2.Error()}
		response := helpers.BaseCommandResponse("error", http.StatusBadRequest, "error", errormessage)

		c.JSON(http.StatusBadRequest, response)
	}
	response := helpers.BaseCommandResponse("Successfully Changed", http.StatusOK, "success", asd)
	c.JSON(http.StatusOK, response)
}
