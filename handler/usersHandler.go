package handler

import (
	"fmt"
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
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	token, err := ch.authService.GenerateToken(users.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	memberDTO := dto.FormatUsersResponse(users, token)
	errsMessage := fmt.Sprintf("Register success%v!", memberDTO)
	response := errs.SuccessRequest(errsMessage)
	c.JSON(http.StatusOK, response)
}

func (ch *UsersHandler) LoginUsers(c *gin.Context) {
	var input domain.Login
	err := c.ShouldBindJSON(&input)
	users, _ := ch.service.LoginUsers(input)
	if err != nil {
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
	memberDTO := dto.FormatUsersResponse(users, token)
	errsMessage := "Login Success"
	response := errs.SuccessRequestUser(errsMessage, memberDTO)
	c.JSON(http.StatusOK, response)
}
