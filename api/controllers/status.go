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

//Status controller -> struct
type StatusController struct {
	logger        infrastructure.Logger
	StatusService services.StatusService
}

//NewStatusController -> constuctor
func NewStatusController(
	logger infrastructure.Logger,
	StatusService services.StatusService,
) StatusController {
	return StatusController{
		logger:        logger,
		StatusService: StatusService,
	}
}

//Create Status
func (cc StatusController) CreateStatus(c *gin.Context) {
	status := models.Status{}

	if err := c.ShouldBindJSON(&status); err != nil {
		err := errors.BadRequest.Wrap(err, "Failed to create status")
		responses.HandleError(c, err)
		return
	}

	_, err := cc.StatusService.CreateStatus(status)

	if err != nil {
		responses.HandleError(c, err)
	}

	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "Status created successfully", "data": status})
}

//GetAllStatus
func (cc StatusController) GetAllStatus(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	pagination.Sort = "created_at_desc"
	priorities, count, err := cc.StatusService.GetAllStatus(pagination)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to find status")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, priorities, count)
}

//DeleteOneStatus
func (cc StatusController) DeleteOneStatus(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := cc.StatusService.DeleteOneStatus(ID)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to delete status")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Status Deleted Successfully")
}
