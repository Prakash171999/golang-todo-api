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
