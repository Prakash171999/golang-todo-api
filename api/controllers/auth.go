package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

const SecretKey = "secret"

type UserAuthController struct {
	logger          infrastructure.Logger
	UserAuthService services.UserAuthService
	jwt             services.JWTService
}

func NewUserAuthController(
	logger infrastructure.Logger,
	UserAuthService services.UserAuthService,
	jwt services.JWTService,
) UserAuthController {
	return UserAuthController{
		logger:          logger,
		UserAuthService: UserAuthService,
		jwt:             jwt,
	}
}

func (cc UserAuthController) CreateUser(c *gin.Context) {
	user := models.User{}

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequest.Wrap(err, "Failed to create user")
		responses.HandleError(c, err)
		return
	}

	registeringUser, _ := cc.UserAuthService.GetUserFromEmail(user.Email)

	//Checking if the user exists
	if registeringUser.Email == user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "User already registered with this email. Please try a different email.",
		})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)

	register_user := models.User{
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    string(password),
	}

	_, err := cc.UserAuthService.CreateUser(register_user)

	if err != nil {
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "User created successfully")
}

func (cc UserAuthController) Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		return
	}

	loggedInUser, err := cc.UserAuthService.GetUserFromEmail(user.Email)

	err1 := bcrypt.CompareHashAndPassword([]byte(loggedInUser.Password), []byte(user.Password))

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Incorrect password",
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "User not registered.",
		})
		return
	}

	token := cc.jwt.GenerateToken(user.Email, true)
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"token":  token,
	})
}

func (cc UserAuthController) ResetPassword(c *gin.Context) {
	user := models.ResetUser{}

	if err := c.ShouldBindJSON(&user); err != nil {
		err := errors.BadRequest.Wrap(err, "failed to reset password")
		responses.HandleError(c, err)
		return
	}

	if user.Password == user.NewPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Old password and new password cannot be same",
		})
		return
	}

	existingUser, err := cc.UserAuthService.GetUserFromEmail(user.Email)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed To Find user")
		responses.HandleError(c, err)
		return
	}

	err1 := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))

	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Incorrect password",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.NewPassword), 14)
	user.Password = string(hashedPassword)

	resetPwdErr := cc.UserAuthService.ResetPassword(user, user.Password)

	if resetPwdErr != nil {
		err := errors.InternalError.Wrap(resetPwdErr, "failed to reset password")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "Password reset successfully"})

}
