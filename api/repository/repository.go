package repository

import "go.uber.org/fx"

// Module exports dependency
var Module = fx.Options(
	fx.Provide(NewTodoRepository),
	fx.Provide(NewPriorityRepository),
	fx.Provide(NewStatusRepository),
	fx.Provide(NewCategoryRepository),
	fx.Provide(NewUserAuthRepository),
	fx.Provide(NewFavouriteRepository),
)
