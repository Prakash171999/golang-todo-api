package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewTodoRoutes),
	fx.Provide(NewPriorityRoutes),
	fx.Provide(NewStatusRoutes),
	fx.Provide(NewCategoryRoutes),
	fx.Provide(NewUserAuthRoutes),
	fx.Provide(NewFavouriteRoute),
	fx.Provide(NewUserRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	todoRoutes TodoRoutes,
	priorityRoutes PriorityRoutes,
	statusRoutes StatusRoutes,
	categoryRoutes CategoryRoutes,
	userAuthRoutes UserAuthRoutes,
	favouriteRoutes FavouriteRoute,
	userRoutes UserRoutes,
) Routes {
	return Routes{
		todoRoutes,
		priorityRoutes,
		statusRoutes,
		categoryRoutes,
		userAuthRoutes,
		favouriteRoutes,
		userRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
