package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
)

type UserAuthRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewUserAuthRepository(db infrastructure.Database, logger infrastructure.Logger) UserAuthRepository {
	return UserAuthRepository{
		db:     db,
		logger: logger,
	}
}

func (c UserAuthRepository) Register(User models.User) (models.User, error) {
	return User, c.db.DB.Create(&User).Error
}
