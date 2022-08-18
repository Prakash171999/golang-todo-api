package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewFirebaseService),
	fx.Provide(NewTwilioService),
	fx.Provide(NewGmailService),
	fx.Provide(NewTodoService),
	fx.Provide(NewPriorityService),
	fx.Provide(NewStatusService),
	fx.Provide(NewCategoryService),
	fx.Provide(NewUserAuthService),
	fx.Provide(NewJWTAuthService),
)
