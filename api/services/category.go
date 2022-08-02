package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

//CategoryService => struct
type CategoryService struct {
	repository repository.CategoryRepository
}

//NewCategoryService -> creates a new Categoryservice
func NewCategoryService(repository repository.CategoryRepository) CategoryService {
	return CategoryService{
		repository: repository,
	}
}

//CreateCategory -> call to create the category
func (c CategoryService) CreateCategory(category models.Category) (models.Category, error) {
	return c.repository.Create(category)
}

//GetAllCategory
func (c CategoryService) GetAllCategory(pagination utils.Pagination) ([]models.Category, int64, error) {
	return c.repository.GetAllCategory(pagination)
}

//DeleteOneCategory
func (c CategoryService) DeleteOneCategory(ID int64) error {
	return c.repository.DeleteOneCategory(ID)
}
