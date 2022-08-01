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

//Priority controller -> struct
type PriorityController struct {
	logger          infrastructure.Logger
	PriorityService services.PriorityService
}

//NewPriorityController -> constuctor
func NewPriorityController(
	logger infrastructure.Logger,
	PriorityService services.PriorityService,
) PriorityController {
	return PriorityController{
		logger:          logger,
		PriorityService: PriorityService,
	}
}

//Create Priority
func (cc PriorityController) CreatePriority(c *gin.Context) {
	priority := models.Priority{}

	if err := c.ShouldBindJSON(&priority); err != nil {
		err := errors.BadRequest.Wrap(err, "Failed to create priority")
		responses.HandleError(c, err)
		return
	}

	_, err := cc.PriorityService.CreatePriority(priority)

	if err != nil {
		responses.HandleError(c, err)
	}

	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "Priority created successfully", "priority": priority})
}

//GetAllPriority
func (cc PriorityController) GetAllPriority(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	pagination.Sort = "created_at_desc"
	priorities, count, err := cc.PriorityService.GetAllPriority(pagination)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to find priority")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, priorities, count)
}

//DeleteOnePriority
func (cc PriorityController) DeleteOnePriority(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := cc.PriorityService.DeleteOnePriority(ID)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to delete priority")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Priority Deleted Successfully")
}
