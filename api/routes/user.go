package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type UserRoutes struct {
	logger         infrastructure.Logger
	router         infrastructure.Router
	userController controllers.UserController
	middleware     middlewares.JWTAuthMiddleware
}

func NewUserRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	userController controllers.UserController,
	middleware middlewares.JWTAuthMiddleware,
) UserRoutes {
	return UserRoutes{
		router:         router,
		logger:         logger,
		userController: userController,
		middleware:     middleware,
	}
}

func (c UserRoutes) Setup() {
	c.logger.Zap.Info(" Setting up user routes")
	users := c.router.Gin.Group("/users").Use(c.middleware.AuthorizeJWT())
	{
		users.GET("", c.userController.GetAllUsers)
		users.GET("/:id", c.userController.GetOneUser)
		//todo.PUT("/:id", c.todoController.UpdateOneTodo)
		//todo.DELETE("/:id", c.todoController.DeleteOneTodo)
	}
}
