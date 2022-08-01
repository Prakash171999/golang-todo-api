package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

//PriorityService => struct
type PriorityService struct {
	repository repository.PriorityRepository
}

//NewPriorityService -> creates a new Priorityservice
func NewPriorityService(repository repository.PriorityRepository) PriorityService {
	return PriorityService{
		repository: repository,
	}
}

//CreatePriority -> call to create the priority
func (c PriorityService) CreatePriority(priority models.Priority) (models.Priority, error) {
	return c.repository.Create(priority)
}

//GetAllPriority
func (c PriorityService) GetAllPriority(pagination utils.Pagination) ([]models.Priority, int64, error) {
	return c.repository.GetAllPriority(pagination)
}

//DeleteOnePriority
func (c PriorityService) DeleteOnePriority(ID int64) error {
	return c.repository.DeleteOnePriority(ID)
}
