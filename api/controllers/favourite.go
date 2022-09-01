package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavouriteController struct {
	logger           infrastructure.Logger
	FavouriteService services.FavouriteService
}

func NewFavouriteController(
	logger infrastructure.Logger,
	FavouriteService services.FavouriteService,
) FavouriteController {
	return FavouriteController{
		logger:           logger,
		FavouriteService: FavouriteService,
	}
}

func (cc FavouriteController) CreateFavourite(c *gin.Context) {
	favourite := models.Favourite{}

	if err := c.ShouldBindJSON(&favourite); err != nil {
		cc.logger.Zap.Error("Error [CreateFavourite] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to create favourite")
		responses.HandleError(c, err)
		return
	}
	_, err := cc.FavouriteService.CreateFavourite(favourite)
	if err != nil {
		responses.HandleError(c, err)
		return
	}
	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "Favourite todo added Successfully", "data": favourite})
}

func (cc FavouriteController) GetAllFavourites(c *gin.Context) {
	pagination := utils.BuildPagination(c)
	pagination.Sort = "created_at_desc"
	favourites, count, err := cc.FavouriteService.GetAllFavourites(pagination)

	if err != nil {
		err := errors.InternalError.Wrap(err, "Failed to find favourites")
		responses.HandleError(c, err)
		return
	}

	responses.JSONCount(c, http.StatusOK, favourites, count)
}
