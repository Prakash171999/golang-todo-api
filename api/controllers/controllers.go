package controllers

import "go.uber.org/fx"

// Module exported for initializing application
var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewTodoController),
	fx.Provide(NewPriorityController),
	fx.Provide(NewStatusController),
	fx.Provide(NewCategoryController),
)
