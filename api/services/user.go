package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) UserService {
	return UserService{
		repository: repository,
	}
}

func (c UserService) GetAllUsers(pagination utils.Pagination) ([]models.User, int64, error) {
	return c.repository.GetAllUsers(pagination)
}
