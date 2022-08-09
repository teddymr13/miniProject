package handler

import (
	"invertory/auth"
	"invertory/domain"
	"invertory/dto"
	"invertory/errs"
	"invertory/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service     service.UsersService
	authService auth.Service
}

func NewUsersHandler(service service.UsersService, authService auth.Service) *UsersHandler {

	return &UsersHandler{service: service, authService: authService}

}

func (ch *UsersHandler) CreateUsers(c *gin.Context) {
	var input domain.UsersInput
	err := c.ShouldBindJSON(&input)
	users, _ := ch.service.CreateUsers(input)
	if err != nil {
		res := errs.NewBadRequestError("Register failed")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, err := ch.authService.GenerateToken(users.ID)
	if err != nil {
		res := errs.NewBadRequestError("Register failed")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	register := dto.FormatUsersResponse(users, token)
	errsMessage := "Register Success"

	response := errs.SuccessRequestUser(errsMessage, register)
	c.JSON(http.StatusOK, response)
}

func (ch *UsersHandler) LoginUsers(c *gin.Context) {
	var input domain.Login
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := errs.NewBadRequestError("Login failed")
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}
	users, errUser := ch.service.LoginUsers(input)
	if errUser != nil {
		res := errs.NewBadRequestError("Login failed")
		c.JSON(http.StatusBadRequest, res)
		return
	}
	token, err := ch.authService.GenerateToken(users.ID)
	if err != nil {
		response := errs.NewBadRequestError("Login failed")
		c.JSON(http.StatusBadRequest, response)
		return
	}
	login := dto.FormatUsersResponse(users, token)
	errsMessage := "Login Success"

	response := errs.SuccessRequestUser(errsMessage, login)
	c.JSON(http.StatusOK, response)
}
