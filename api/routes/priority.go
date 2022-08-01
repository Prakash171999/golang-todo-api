package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

//PriorityRoutes -> struct
type PriorityRoutes struct {
	logger             infrastructure.Logger
	router             infrastructure.Router
	priorityController controllers.PriorityController
	middleware         middlewares.FirebaseAuthMiddleware
}

//NewPriorityRoutes -> creates new Priority routes
func NewPriorityRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	priorityController controllers.PriorityController,
	middleware middlewares.FirebaseAuthMiddleware,
) PriorityRoutes {
	return PriorityRoutes{
		router:             router,
		logger:             logger,
		priorityController: priorityController,
		middleware:         middleware,
	}
}

//Setup priority routes
func (c PriorityRoutes) Setup() {
	priority := c.router.Gin.Group("/priorities")
	{
		priority.POST("", c.priorityController.CreatePriority)
		priority.GET("", c.priorityController.GetAllPriority)
		priority.DELETE("/:id", c.priorityController.DeleteOnePriority)
	}
}
