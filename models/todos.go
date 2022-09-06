package models

type Todo struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	DueDate     string `json:"due_date"`
	Image       string `json:"image"`
	StatusId    *int   `json:"statusId" gorm:"column:statusId"`
	PriorityId  *int   `json:"priorityId" gorm:"column:priorityId"`
	CategoryId  *int   `json:"categoryId" gorm:"column:categoryId"`
}

func (m Todo) TableName() string {
	return "todos"
}

func (m Todo) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"title":       m.Title,
		"description": m.Description,
		"image":       m.Image,
		"start_date":  m.StartDate,
		"due_date":    m.DueDate,
		"statusId":    m.StatusId,
		"priorityId":  m.PriorityId,
		"categoryId":  m.CategoryId,
	}
}
