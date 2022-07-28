package routes

import (
    "boilerplate-api/api/controllers"
    "boilerplate-api/api/middlewares"
    "boilerplate-api/infrastructure"
)


// TodoRoutes -> struct 
type TodoRoutes Struct{
	logger                    infrastructure.Logger
    router                    infrastructure.Router
    todoController controllers.TodoController
    middleware                middlewares.FirebaseAuthMiddleware
}

// NewTodoRoutes -> creates new Todo controller
func NewTodoRoutes(logger infrastructure.Logger,
    router infrastructure.Router,
    todoController controllers.TodoController,
    middleware middlewares.FirebaseAuthMiddleware,) TodoRoutes {
		return TodoRoutes{
			router:                    router,
			logger:                    logger,
			todoController: todoController,
			middleware:                middleware,
		}
	}

// Setup todo routes
func (c TodoRoutes) Setup() {
	c.logger.Zap.Info(" Setting up Todo routes")
    todo := c.router.Gin.Group("/todos")
	{
		todo.POST("todo", c.todoController.CreateTodo)
		todo.GET("todo/all",c.todoController.GetAllTodo)
		todo.GET("todo/:id", c.todoController.UpdateOneTodo)
		todo.Delete("todo/delete/:id", c.todoController.DeleteOneTodo)
	}
}