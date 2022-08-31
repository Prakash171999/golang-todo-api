package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

type FavouriteService struct {
	repository repository.FavouriteRepository
}

func NewFavouriteService(repository repository.FavouriteRepository) FavouriteService {
	return FavouriteService{
		repository: repository,
	}
}

func (c FavouriteService) CreateFavourite(favourite models.Favourite) (models.Favourite, error) {
	return c.repository.Create(favourite)
}

func (c FavouriteService) GetAllFavourites(pagination utils.Pagination) ([]models.Favourite, int64, error) {
	return c.repository.GetAllFavourites(pagination)
}
