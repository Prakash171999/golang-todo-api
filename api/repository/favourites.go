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

func (c FavouriteRepository) GetAllFavourites(pagination utils.Pagination) (favourites []models.UserFavourites, count int64, err error) {
	queryBuider := c.db.DB.Model(&models.Todo{}).
		Select(
			`users.id AS userId,
			  todos.id AS todoId,
			  todos.priorityId AS priority_id,
			  users.full_name AS user_name,
			  users.email AS email,
			  users.phone_number AS phone_no,
			  todos.title AS todo_title,
			  todos.description AS todo_description,
			  todos.image AS todo_image,
			  priorities.priority_type AS priority,
			  status.status_type AS status`).Offset(pagination.Offset)

	if !pagination.All {
		queryBuider = queryBuider.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuider.Where(c.db.DB.Where("`favourite`.`id` LIKE ?", searchQuery))
	}

	return favourites, count, queryBuider.Joins("INNER JOIN favourites on todos.id=favourites.todoId").
		Joins("LEFT JOIN users on users.id = favourites.userId").
		Joins("LEFT JOIN priorities on priorities.id = todos.priorityId").
		Joins("LEFT JOIN status on status.id = todos.statusId").
		Find(&favourites).
		Offset(-1).
		Limit(-1).
		Count(&count).Error
}

func (c FavouriteRepository) DeleteUserFavourite(ID int64) error {
	return c.db.DB.Where("id = ?", ID).Delete(&models.Favourite{}).Error
}
