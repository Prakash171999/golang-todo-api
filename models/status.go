package models

type Status struct {
	Base
	StatusType string `json:"status_type"`
}

func (m Status) TableName() string {
	return "status"
}

func (m Status) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"status_type": m.StatusType,
	}
}
