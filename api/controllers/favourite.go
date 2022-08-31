package controllers

import (
	"boilerplate-api/api/responses"
	"boilerplate-api/api/services"
	"boilerplate-api/errors"
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
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
		cc.logger.Zap.Error("Error [CreateTodo] (ShouldBindJson) : ", err)
		err := errors.BadRequest.Wrap(err, "Failed to create todo")
		responses.HandleError(c, err)
		return
	}
	_, err := cc.FavouriteService.CreateFavourite(favourite)
	if err != nil {
		responses.HandleError(c, err)
	}
	responses.SuccessJSON(c, http.StatusOK, gin.H{"status": "Favourite todo added Successfully", "data": favourite})
}
