package controllers

import (
	"strconv"
	"net/http"
    "boilerplate-api/api/responses"
    "boilerplate-api/api/services"
    "boilerplate-api/infrastructure"
    "boilerplate-api/models"
    "boilerplate-api/utils"
    "boilerplate-api/errors"

    "github.com/gin-gonic/gin"
	"fmt"

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

// GetOneTodo -> Get One Todo
func (cc TodoController) GetOneTodo(c *gin.Context) {
    ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    todo, err := cc.TodoService.GetOneTodo(ID)

    if err != nil {
        cc.logger.Zap.Error("Error [GetOneTodo] [db GetOneTodo]: ", err.Error())
        err := errors.InternalError.Wrap(err, "Failed To Find Todo")
        responses.HandleError(c, err)
        return
    }

	fmt.Println("todo", todo.StatusId)

    responses.JSON(c, http.StatusOK, &todo)

}

// // UpdateOneTodo -> Update One Todo By Id
func (cc TodoController) UpdateOneTodo(c *gin.Context) {
    ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
    todo := models.Todo{}

    if err := c.ShouldBindJSON(&todo); err != nil {
        cc.logger.Zap.Error("Error [UpdateTodo] (ShouldBindJson) : ", err)
        err := errors.BadRequest.Wrap(err, "failed to update todo")
        responses.HandleError(c, err)
        return
    }
    todo.ID = ID

	updateTodo, err := cc.TodoService.GetOneTodo(ID)

    if err != nil {
        cc.logger.Zap.Error("Error [GetOneTodo] [db GetOneTodo]: ", err.Error())
        err := errors.InternalError.Wrap(err, "Failed To Find Todo")
        responses.HandleError(c, err)
        return
    }

	updateTodo.StatusId = todo.StatusId

    if err := cc.TodoService.UpdateOneTodo(*updateTodo); err != nil {
        cc.logger.Zap.Error("Error [UpdateTodo] [db UpdateTodo]: ", err.Error())
        err := errors.InternalError.Wrap(err, "failed to update todo")
        responses.HandleError(c, err)
        return
    }

    responses.SuccessJSON(c, http.StatusOK, "Updated successfully")
}
