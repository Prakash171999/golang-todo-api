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
) Routes {
	return Routes{
		todoRoutes,
		priorityRoutes,
		statusRoutes,
		categoryRoutes,
		userAuthRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
