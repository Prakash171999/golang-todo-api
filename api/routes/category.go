package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type CategoryRoutes struct {
	logger           infrastructure.Logger
	router           infrastructure.Router
	statusController controllers.CategoryController
	middleware       middlewares.FirebaseAuthMiddleware
}

func NewCategoryRoutes(
	logger infrastructure.Logger,
	router infrastructure.Router,
	statusController controllers.CategoryController,
	middleware middlewares.FirebaseAuthMiddleware,
) CategoryRoutes {
	return CategoryRoutes{
		router:           router,
		logger:           logger,
		statusController: statusController,
		middleware:       middleware,
	}
}

func (c CategoryRoutes) Setup() {
	category := c.router.Gin.Group("/category")
	{
		category.POST("", c.statusController.CreateCategory)
		category.GET("", c.statusController.GetAllCategory)
		category.DELETE("/:id", c.statusController.DeleteOneCategory)
	}
}
