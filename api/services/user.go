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

func (c UserService) GetOneUser(ID int64) (*models.User, error) {
	return c.repository.GetOneUser(ID)
}

func (c UserService) UpdateOneUser(user models.User) (models.User, error) {
	return c.repository.UpdateOneUser(user)
}

func (c UserService) DeleteOneUser(ID int64) error {
	return c.repository.DeleteOneUser(ID)
}
