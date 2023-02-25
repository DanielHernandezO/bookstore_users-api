package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/DanielHernandezO/bookstore_users-api/domain"
	"github.com/DanielHernandezO/bookstore_users-api/service"
	"github.com/gin-gonic/gin"
)

type UserController interface {
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) *userController {
	return &userController{
		userService: userService,
	}
}

func (u userController) GetUser(c *gin.Context) {
	var user domain.User
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		//handle the error --> to do
		return
	}
	if err := json.Unmarshal(bytes, &user); err != nil {
		// handler the error  --> to do
		return
	}
	result, err := u.userService.CreateUser(user)
	if err != nil {
		// handler the error  --> to do
		return
	}
	c.JSON(http.StatusCreated, result)
}

func (u userController) CreateUser(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Implement me!")
}
