package models

type Favourite struct {
	Base
	UserId *int `json:"userId" gorm:"column:userId"`
	TodoId *int `json:"todoId" gorm:"column:todoId"`
}

func (m Favourite) TableName() string {
	return "favourites"
}

func (m Favourite) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"userId": m.UserId,
		"todoId": m.TodoId,
	}
}
