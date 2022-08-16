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

	// var existingUser models.User

	// emailExists := c.db.DB.Where("email = ?", User.Email).First(&existingUser)

	// if emailExists.RowsAffected > 0 {
	// 	fmt.Println("Email exists. Please use another email address.")
	// 	return existingUser, c.db.DB.Error
	// }

	return User, c.db.DB.Create(&User).Error

}
