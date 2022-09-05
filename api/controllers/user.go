package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	logger      infrastructure.Logger
	UserService services.UserService
}

func NewUserController(
	logger infrastructure.Logger,
	UserService services.UserService,
) UserController {
	return UserController{
		logger:      logger,
		UserService: UserService,
	}
}

func (cc UserController) GetAllUsers(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	pagination.Sort = "created_at_desc"
	users, count, err := cc.UserService.GetAllUsers(pagination)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to find users")
		responses.HandleError(c, err)
		return
	}
	responses.JSONCount(c, http.StatusOK, users, count)
}

func (cc UserController) GetOneUser(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user, err := cc.UserService.GetOneUser(ID)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to find the user")
		responses.HandleError(c, err)
		return
	}

	responses.JSON(c, http.StatusOK, &user)
}

func (cc UserController) UpdateOneUser(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequest.Wrap(err, "Failed to update the user")
		responses.HandleError(c, err)
		return
	}

	updateUser, err := cc.UserService.GetOneUser(ID)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed To Find the user")
		responses.HandleError(c, err)
		return
	}

	user.ID = updateUser.ID
	updatedUser, err := cc.UserService.UpdateOneUser(user)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to update user")
		responses.HandleError(c, err)
		responses.JSON(c, http.StatusOK, updateUser)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "User updated successfully", "updatedData": updatedUser})

}
