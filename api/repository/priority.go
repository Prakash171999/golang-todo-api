package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

//TodoRepository database structure
type PriorityRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

//NewPriorityRepository -> creates a new Priority repository
func NewPriorityRepository(db infrastructure.Database, logger infrastructure.Logger) PriorityRepository {
	return PriorityRepository{
		db:     db,
		logger: logger,
	}
}

//create priority
func (c PriorityRepository) Create(Priority models.Priority) (models.Priority, error) {
	return Priority, c.db.DB.Create(&Priority).Error
}

//GetAllPriority
func (c PriorityRepository) GetAllPriority(pagination utils.Pagination) ([]models.Priority, int64, error) {
	var priorities []models.Priority
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Model(&models.Todo{}).Offset(pagination.Offset)

	if !pagination.All {
		queryBuilder = queryBuilder.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuilder.Where(c.db.DB.Where("`priorities`. `title` LIKE ?", searchQuery))
	}

	err := queryBuilder.Find(&priorities).Offset(-1).Limit(-1).Count(&totalRows).Error
	return priorities, totalRows, err
}

//DeleteOnePriority
func (c PriorityRepository) DeleteOnePriority(ID int64) error {
	return c.db.DB.Where("id = ?", ID).Delete(&models.Priority{}).Error
}
