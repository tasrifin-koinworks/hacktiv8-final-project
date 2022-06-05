package controllers

import (
	"hacktiv8-final-project/params"
	"hacktiv8-final-project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{
		userService: *service,
	}
}

func (u *UserController) UserRegister(c *gin.Context) {
	var req params.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.UserRegister(req)

	c.JSON(result.Status, result.Payload)
}

func (u *UserController) Login(c *gin.Context) {
	var req params.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	result := u.userService.Login(req)

	c.JSON(result.Status, result.Payload)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var req params.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	id, _ := strconv.Atoi(c.GetString("id"))
	userId, _ := strconv.Atoi(c.Param("userId"))

	req.ID = userId
	result := u.userService.UpdateUser(id, req)
	c.JSON(result.Status, result.Payload)

}

func (u *UserController) DeleteUser(c *gin.Context) {
	var req params.CreateUser

	err := c.ShouldBind(&req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Payload: err.Error(),
		})

		return
	}

	id, _ := strconv.Atoi(c.GetString("id"))
	userId, _ := strconv.Atoi(c.Param("userId"))

	req.ID = userId
	result := u.userService.DeleteUser(id, req)
	c.JSON(result.Status, result.Payload)
}
