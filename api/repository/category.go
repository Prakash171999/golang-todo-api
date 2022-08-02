package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

type CategoryRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewCategoryRepository(db infrastructure.Database, logger infrastructure.Logger) CategoryRepository {
	return CategoryRepository{
		db:     db,
		logger: logger,
	}
}

func (c CategoryRepository) Create(Category models.Category) (models.Category, error) {
	return Category, c.db.DB.Create(&Category).Error
}

func (c CategoryRepository) GetAllCategory(pagination utils.Pagination) ([]models.Category, int64, error) {
	var priorities []models.Category
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Model(&models.Category{}).Offset(pagination.Offset)

	if !pagination.All {
		queryBuilder = queryBuilder.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`categories`. `title` LIKE ?", searchQuery))
	}

	err := queryBuilder.Find(&priorities).Offset(-1).Limit(-1).Count(&totalRows).Error
	return priorities, totalRows, err
}

func (c CategoryRepository) DeleteOneCategory(ID int64) error {
	return c.db.DB.Where("id = ?", ID).Delete(&models.Category{}).Error
}
