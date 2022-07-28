package models

type Category struct {
    Base
    category_type          string `json:"category_type"` 
}

func (m Category) TableName() string {
    return "categories"
}

func (m Category) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "category_type":           m.category_type,
}
}