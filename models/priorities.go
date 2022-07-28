package models

type Priority struct {
    Base
    PriorityType          string `json:"priority_type"` 
}

func (m Priority) TableName() string {
    return "priorities"
}

func (m Priority) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "priority_type":           m.PriorityType,
}
}