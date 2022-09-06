package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
	"boilerplate-api/utils"
	"fmt"
)

// TodoRepository database structure
type TodoRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

// NewTodoRepository creates a new Todo repository
func NewTodoRepository(db infrastructure.Database, logger infrastructure.Logger) TodoRepository {
	return TodoRepository{
		db:     db,
		logger: logger,
	}
}

// Create Todo
func (c TodoRepository) Create(Todo models.Todo) (models.Todo, error) {
	return Todo, c.db.DB.Create(&Todo).Error
}

// GetAllTodo -> Get All todos
func (c TodoRepository) GetAllTodo(pagination utils.Pagination) ([]models.Todo, int64, error) {
	var todos []models.Todo
	var totalRows int64 = 0
	queryBuider := c.db.DB.Model(&models.Todo{}).Offset(pagination.Offset)

	if !pagination.All {
		queryBuider = queryBuider.Limit(pagination.PageSize)
	}

	if pagination.Keyword != "" {
		searchQuery := "%" + pagination.Keyword + "%"
		queryBuider.Where(c.db.DB.Where("`todos`.`title` LIKE ?", searchQuery))
	}

	err := queryBuider.
		Find(&todos).
		Offset(-1).
		Limit(-1).
		Count(&totalRows).Error
	return todos, totalRows, err
}

// GetOneTodo -> Get One Todo By Id
func (c TodoRepository) GetOneTodo(ID int64) (*models.Todo, error) {
	todo := models.Todo{}
	// return &Todo, c.db.DB.
	//     Where("id = ?", ID).First(&Todo).Error
	err := c.db.DB.Where("id=?", ID).First(&todo).Error
	if err != nil {
		fmt.Println(err)
	}
	return &todo, err
}

// UpdateOneTodo -> Update One Todo By Id
func (c TodoRepository) UpdateOneTodo(todo models.Todo) (models.Todo, error) {
	return todo, c.db.DB.Model(&models.Todo{}).
		Where("id = ?", todo.ID).
		Updates(map[string]interface{}{
			"title":       todo.Title,
			"description": todo.Description,
			"image":       todo.Image,
			"start_date":  todo.StartDate,
			"end_date":    todo.DueDate,
			"statusId":    todo.StatusId,
			"priorityId":  todo.PriorityId,
			"categoryId":  todo.CategoryId,
		}).Find(&todo).Error
}

// DeleteOneTodo -> Delete One Todo By Id
func (c TodoRepository) DeleteOneTodo(ID int64) error {
	return c.db.DB.
		Where("id = ?", ID).
		Delete(&models.Todo{}).
		Error
}
