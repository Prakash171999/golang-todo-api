package models

type Statu struct {
    Base
    StatusType          string `json:"status_type"` 
}

func (m Statu) TableName() string {
    return "status"
}

func (m Statu) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "status_type":           m.StatusType,
}
}