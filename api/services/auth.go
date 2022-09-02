package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
)

type UserAuthService struct {
	repository repository.UserAuthRepository
}

func NewUserAuthService(repository repository.UserAuthRepository) UserAuthService {
	return UserAuthService{
		repository: repository,
	}
}

func (c UserAuthService) CreateUser(user models.User) (models.User, error) {
	return c.repository.Register(user)
}

func (c UserAuthService) GetUserFromEmail(userEmail string) (models.User, error) {

	return c.repository.GetUserFromEmail(userEmail)
}

func (c UserAuthService) ResetPassword(user models.User, password string) error {
	return c.repository.ResetPassword(user, password)
}

// func (c UserAuthService) LoginUser(user models.User) bool {
// 	// var user models.User
// 	// if err := cc.db.DB.Where("email = ?", email).Where("password = ?", password).First(&user).Error; err != nil {
// 	// 	return false
// 	// }
// 	// return true
// 	return c.repository.Login(user)
// }
