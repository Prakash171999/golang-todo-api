package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

type FavouriteRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewFavouriteRepository(db infrastructure.Database, logger infrastructure.Logger) FavouriteRepository {
	return FavouriteRepository{
		db:     db,
		logger: logger,
	}
}

func (c FavouriteRepository) Create(Favourite models.Favourite) (models.Favourite, error) {
	return Favourite, c.db.DB.Create(&Favourite).Error
}

func (c FavouriteRepository) GetAllFavourites(pagination utils.Pagination) ([]models.Favourite, int64, error) {
	var favourite []models.Favourite
	var totalRows int64 = 0
	queryBuider := c.db.DB.Model(&models.Favourite{}).Offset(pagination.Offset)

	if !pagination.All {
		queryBuider = queryBuider.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuider.Where(c.db.DB.Where("`favourite`.`id` LIKE ?", searchQuery))
	}

	err := queryBuider.
		Find(&favourite).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return favourite, totalRows, err
}
