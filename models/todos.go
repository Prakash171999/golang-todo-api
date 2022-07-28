package models

type Todo struct {
    Base
    Title          string `json:"title"`
    Description       string `json:"description"`
    Image           string `json:"image"`   
}

func (m Todo) TableName() string {
    return "todos"
}

func (m Todo) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "title":           m.Title,
        "description":       m.Description,
        "image":       m.Image,
    }
}