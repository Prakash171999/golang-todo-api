package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TodoController -> struct
type TodoController struct {
	logger      infrastructure.Logger
	TodoService services.TodoService
}

// NewTodoController -> constructor
func NewTodoController(
	logger infrastructure.Logger,
	TodoService services.TodoService,
) TodoController {
	return TodoController{
		logger:      logger,
		TodoService: TodoService,
	}
}

//CreateTodo -> Create single todo
func (cc TodoController) CreateTodo(c *gin.Context) {
	todo := models.Todo{}

	if err := c.ShouldBindJSON(&todo); err != nil {
		cc.logger.Zap.Error("Error [CreateTodo] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to create todo")
		responses.HandleError(c, err)
		return
	}
	_, err := cc.TodoService.CreateTodo(todo)
	if err != nil {
		responses.HandleError(c, err)
	}
	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "Todo Created Successfully", "data": todo})
}

//GetAllTodo -> Get All Todo list
func (cc TodoController) GetAllTodo(c *gin.Context) {
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

	updateTodo, err := cc.TodoService.GetOneTodo(ID)

	if err != nil {
		cc.logger.Zap.Error("Error [GetOneTodo] [db GetOneTodo]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed To Find Todo")
		responses.HandleError(c, err)
		return
	}

	todo.ID = updateTodo.ID

	// updateTodo.Title = todo.Title
	// updateTodo.Description = todo.Description
	// updateTodo.Image = todo.Image
	// updateTodo.StatusId = todo.StatusId
	// updateTodo.PriorityId = todo.PriorityId
	// updateTodo.CategoryId = todo.CategoryId
	updatedTodo, err := cc.TodoService.UpdateOneTodo(todo)

	if err != nil {
		cc.logger.Zap.Error("Error [UpdateTodo] [db UpdateTodo]: ", err.Error())
		err := errors.InternalError.Wrap(err, "failed to update todo")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "Todo updated successfully", "updatedData": updatedTodo})
}

// DeleteOneTodo -> Delete One Todo By Id
func (cc TodoController) DeleteOneTodo(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := cc.TodoService.DeleteOneTodo(ID)

	if err != nil {
		cc.logger.Zap.Error("Error [DeleteOneTodo] [db DeleteOneTodo]: ", err.Error())
		err := errors.InternalError.Wrap(err, "Failed to Delete Todo")
		responses.HandleError(c, err)
		return
	}

	responses.SuccessJSON(c, http.StatusOK, "Todo Deleted Sucessfully")
}
