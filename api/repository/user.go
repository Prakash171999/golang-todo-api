package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"fmt"
)

type UserRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewUserRepository(db infrastructure.Database, logger infrastructure.Logger) UserRepository {
	return UserRepository{
		db:     db,
		logger: logger,
	}
}

func (c UserRepository) GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	var users []models.User
	var totalRows int64 = 0
	queryBuider := c.db.DB.Model(&models.User{}).Offset(pagination.Offset)

	if !pagination.All {
		queryBuider = queryBuider.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuider.Where(c.db.DB.Where("`User`.`full_name` LIKE ?", searchQuery))
	}

	err := queryBuider.
		Find(&users).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return users, totalRows, err
}

func (c UserRepository) GetOneUser(ID int64) (*models.User, error) {
	user := models.User{}
	err := c.db.DB.Where("id = ?", ID).First(&user).Error
	if err != nil {
		fmt.Println(err)
	}
	return &user, err
}

func (c UserRepository) UpdateOneUser(user models.User) (models.User, error) {
	return user, c.db.DB.Model(&models.User{}).
		Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"full_name":    user.FullName,
			"phone_number": user.PhoneNumber,
		}).Find(&user).Error
}

func (c UserRepository) DeleteOneUser(ID int64) error {
	return c.db.DB.Where("id = ?", ID).
		Delete(&models.User{}).Error
}
