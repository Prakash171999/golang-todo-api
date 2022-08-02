package services

import (
	"boilerplate-api/api/repository"
	"boilerplate-api/models"
	"boilerplate-api/utils"
)

//StatusService => struct
type StatusService struct {
	repository repository.StatusRepository
}

//NewStatusService -> creates a new Statusservice
func NewStatusService(repository repository.StatusRepository) StatusService {
	return StatusService{
		repository: repository,
	}
}

//CreateStatus -> call to create the status
func (c StatusService) CreateStatus(status models.Status) (models.Status, error) {
	return c.repository.Create(status)
}

//GetAllStatus
func (c StatusService) GetAllStatus(pagination utils.Pagination) ([]models.Status, int64, error) {
	return c.repository.GetAllStatus(pagination)
}

//DeleteOneStatus
func (c StatusService) DeleteOneStatus(ID int64) error {
	return c.repository.DeleteOneStatus(ID)
}
