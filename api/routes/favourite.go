package routes

import (
	"boilerplate-api/api/controllers"
	"boilerplate-api/api/middlewares"
	"boilerplate-api/infrastructure"
)

type FavouriteRoute struct {
	logger              infrastructure.Logger
	router              infrastructure.Router
	favouriteController controllers.FavouriteController
	middleware          middlewares.JWTAuthMiddleware
}

func NewFavouriteRoute(
	logger infrastructure.Logger,
	router infrastructure.Router,
	favouriteController controllers.FavouriteController,
	middleware middlewares.JWTAuthMiddleware,
) FavouriteRoute {
	return FavouriteRoute{
		router:              router,
		logger:              logger,
		favouriteController: favouriteController,
		middleware:          middleware,
	}
}

func (c FavouriteRoute) Setup() {
	favourite := c.router.Gin.Group("/favourite").Use(c.middleware.AuthorizeJWT())
	{
		favourite.POST("", c.favouriteController.CreateFavourite)
		favourite.GET("", c.favouriteController.GetAllFavourites)
		//favourite.DELETE("/:id", c.favouriteController.DeleteOneFavourite)
	}
}
