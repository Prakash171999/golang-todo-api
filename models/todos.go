package models

type Todo struct {
    Base
    Title           string `json:"title"`
    Description     string `json:"description"`
    Image           string `json:"image"`  
    StatusId        *int `json:"statusId" gorm:"column:statusId"`
    PriorityId      *int `json:"priorityId" gorm:"column:priorityId"`
    CategoryId      *int `json:"categoryId" gorm:"column:categoryId"`
}

func (m Todo) TableName() string {
    return "todos"
}

func (m Todo) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "title":           m.Title,
        "description":       m.Description,
        "image":       m.Image,
        "statusId":     m.StatusId,
        "priorityId":   m.PriorityId,
        "categoryId":   m.CategoryId,     
    }
}