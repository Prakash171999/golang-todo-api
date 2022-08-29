package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	fmt.Println("pwwww=>", password, string(password))

	register_user := models.User{
		FullName:    user.FullName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Password:    string(password),
	}

	_, err := cc.UserAuthService.CreateUser(register_user)

	if err != nil {
		responses.HandleError(c, err)
	}

	responses.SuccessJSON(c, http.StatusOK, "User created successfully")
}

func (cc UserAuthController) Login(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		return
	}

	// password := bcrypt.CompareHashAndPassword([]byte(user.Password), user.Password)

	// fmt.Println("cont user pass", password)

	user, err = cc.UserAuthService.GetUserFromEmail(user.Email)

	fmt.Println("user", user)

	if err != nil {
		token := cc.jwt.GenerateToken(user.Email, true)
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"token":  token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  401,
			"message": "User not registered.",
		})
		return
	}
}

// func (cc UserAuthController) LoginUser(c *gin.Context) {
// 	var user_info models.User

// 	if err := c.ShouldBindJSON(&user_info); err != nil {
// 		c.JSON(http.StatusBadRequest, "error found in given data")
// 	}

// 	user_obj, err := services.UserAuthService.GetUserFromEmail(services.UserAuthService{}, user_info.Email)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "error while fetching user from given email")
// 	}
// 	if user_obj.ID == nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"msg": "incorrect email address"})
// 	}

// 	if err := bcrypt.CompareHashAndPassword(user_info.Password, []byte(user_obj.Password)); err != nil {
// 		responses.HandleError(c, err)
// 	}

// 	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
// 		Issuer:    strconv.Itoa(int(*user_info.ID)),
// 		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
// 	})

// 	token, err := claims.SignedString([]byte(SecretKey))

// 	if err != nil {
// 		c.Status(http.StatusInternalServerError)
// 		c.JSON(http.StatusInternalServerError, "Could not login")
// 	}

// 	http.SetCookie(c.Writer, &http.Cookie{
// 		Name:     "refresh",
// 		Value:    token,
// 		Expires:  time.Now().Add(time.Hour * 24),
// 		Secure:   false,
// 		HttpOnly: true, //because frontend doesn't need to access it
// 	})

// 	c.JSON(http.StatusOK, token)
// }
