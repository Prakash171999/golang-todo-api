package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserAuthController struct {
	logger          infrastructure.Logger
	UserAuthService services.UserAuthService
}

func NewUserAuthController(logger infrastructure.Logger, UserAuthService services.UserAuthService) UserAuthController {
	return UserAuthController{
		logger:          logger,
		UserAuthService: UserAuthService,
	}
}

func (cc UserAuthController) CreateUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequest.Wrap(err, "Failed to create user")
		responses.HandleError(c, err)
		return
	}

	_, err := cc.UserAuthService.CreateUser(user)

	if err != nil {
		responses.HandleError(c, err)
	}

	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "User created successfully", "user": user})
}
