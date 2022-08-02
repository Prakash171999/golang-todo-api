package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type StatusRoutes struct {
	logger           infrastructure.Logger
	router           infrastructure.Router
	statusController controllers.StatusController
	middleware       middlewares.FirebaseAuthMiddleware
}

func NewStatusRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	statusController controllers.StatusController,
	middleware middlewares.FirebaseAuthMiddleware,
) StatusRoutes {
	return StatusRoutes{
		router:           router,
		logger:           logger,
		statusController: statusController,
		middleware:       middleware,
	}
}

func (c StatusRoutes) Setup() {
	status := c.router.Gin.Group("/status")
	{
		status.POST("", c.statusController.CreateStatus)
		status.GET("", c.statusController.GetAllStatus)
		status.DELETE("/:id", c.statusController.DeleteOneStatus)
	}
}
