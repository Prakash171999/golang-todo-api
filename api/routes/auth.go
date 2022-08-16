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
	middleware         middlewares.FirebaseAuthMiddleware
}

func NewUserAuthRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	userAuthController controllers.UserAuthController,
	middleware middlewares.FirebaseAuthMiddleware,
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
		user.POST("/login", c.userAuthController.LoginUser)
	}
}
