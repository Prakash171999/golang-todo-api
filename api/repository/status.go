package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

type StatusRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewStatusRepository(db infrastructure.Database, logger infrastructure.Logger) StatusRepository {
	return StatusRepository{
		db:     db,
		logger: logger,
	}
}

func (c StatusRepository) Create(Status models.Status) (models.Status, error) {
	return Status, c.db.DB.Create(&Status).Error
}

func (c StatusRepository) GetAllStatus(pagination utils.Pagination) ([]models.Status, int64, error) {
	var priorities []models.Status
	var totalRows int64 = 0
	queryBuilder := c.db.DB.Model(&models.Status{}).Offset(pagination.Offset)

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

//DeleteOneStatus
func (c StatusRepository) DeleteOneStatus(ID int64) error {
	return c.db.DB.Where("id = ?", ID).Delete(&models.Status{}).Error
}
