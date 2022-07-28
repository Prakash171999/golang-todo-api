package models

type Category struct {
    Base
    userId          string `json:"userId"` 
	todoId          string `json:"userId"` 
}

func (m Category) TableName() string {
    return "categories"
}

func (m Category) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "category_type":           m.CategoryType,
}
}