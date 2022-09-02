package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type UserAuthRoutes struct {
	logger             infrastructure.Logger
	router             infrastructure.Router
	userAuthController controllers.UserAuthController
	middleware         middlewares.JWTAuthMiddleware
}

func NewUserAuthRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	userAuthController controllers.UserAuthController,
	middleware middlewares.JWTAuthMiddleware,
) UserAuthRoutes {
	return UserAuthRoutes{
		router:             router,
		logger:             logger,
		userAuthController: userAuthController,
		middleware:         middleware,
	}
}

func (c UserAuthRoutes) Setup() {
	user := c.router.Gin.Group("/user")
	{
		user.POST("/register", c.userAuthController.CreateUser)
		user.POST("/login", c.userAuthController.Login)
		user.PUT("/reset-password", c.userAuthController.ResetPassword)
	}
}
