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

//Category controller -> struct
type CategoryController struct {
	logger          infrastructure.Logger
	CategoryService services.CategoryService
}

//NewCategoryController -> constuctor
func NewCategoryController(
	logger infrastructure.Logger,
	CategoryService services.CategoryService,
) CategoryController {
	return CategoryController{
		logger:          logger,
		CategoryService: CategoryService,
	}
}

//Create Category
func (cc CategoryController) CreateCategory(c *gin.Context) {
	category := models.Category{}

	if err := c.ShouldBindJSON(&category); err != nil {
		err := errors.BadRequest.Wrap(err, "Failed to create category")
		responses.HandleError(c, err)
		return
	}

	_, err := cc.CategoryService.CreateCategory(category)

	if err != nil {
		responses.HandleError(c, err)
	}

	responses.SuccessJSON(c, http.StatusOK, gin.H{"category": "Category created successfully", "data": category})
}

//GetAllCategory
func (cc CategoryController) GetAllCategory(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	pagination.Sort = "created_at_desc"
	priorities, count, err := cc.CategoryService.GetAllCategory(pagination)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to find category")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, priorities, count)
}

//DeleteOneCategory
func (cc CategoryController) DeleteOneCategory(c *gin.Context) {
	ID, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	err := cc.CategoryService.DeleteOneCategory(ID)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to delete category")
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, "Category Deleted Successfully")
}
