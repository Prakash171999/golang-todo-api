package controllers

import (
	"net/http"
    "boilerplate-api/api/responses"
    "boilerplate-api/api/services"
    "boilerplate-api/infrastructure"
    "boilerplate-api/models"
    "boilerplate-api/utils"
    "boilerplate-api/errors"

    "github.com/gin-gonic/gin"

)

// TodoController -> struct
type TodoController struct {
    logger                 infrastructure.Logger
    TodoService  services.TodoService
}

// NewTodoController -> constructor
func NewTodoController(
    logger infrastructure.Logger,
    TodoService services.TodoService,
) TodoController {
    return TodoController{
        logger:                  logger,
        TodoService:  TodoService,
    }
}

//CreateTodo -> Create single todo
func (cc TodoController) CreateTodo(c *gin.Context){
	todo := models.Todo{}


	if err := c.ShouldBindJSON(&todo); err != nil{
		cc.logger.Zap.Error("Error [CreateTodo] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to create todo")
		responses.HandleError(c, err)
		return
	}
	_, err := cc.TodoService.CreateTodo(todo)
	if err != nil{
		responses.HandleError(c, err)
	}
	responses.SuccessJSON(c, http.StatusOK, "Todo Created Successfully")
}

//GetAllTodo -> Get All Todo list
func (cc TodoController) GetAllTodo(c *gin.Context){
	pagination := utils.BuildPagination(c)
	pagination.Sort = "created_at_desc"
	todos, count, err := cc.TodoService.GetAllTodo(pagination)

	if err != nil {
		cc.logger.Zap.Error("Error finding Todo records", err.Error())
        err := errors.InternalError.Wrap(err, "Failed To Find Todo")
        responses.HandleError(c, err)
        return
	}
	responses.JSONCount(c, http.StatusOK, todos, count)
}