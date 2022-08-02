package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewRoutes),
	fx.Provide(NewTodoRoutes),
	fx.Provide(NewUserRoutes),
	fx.Provide(NewPriorityRoutes),
	fx.Provide(NewStatusRoutes),
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
	userRoutes UserRoutes,
	priorityRoutes PriorityRoutes,
	statusRoutes StatusRoutes,
) Routes {
	return Routes{
		todoRoutes,
		userRoutes,
		priorityRoutes,
		statusRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
