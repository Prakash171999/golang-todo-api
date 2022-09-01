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

type UserFavourites struct {
	UserId          *int   `json:"userId" gorm:"column:userId"`
	TodoId          *int   `json:"todoId" gorm:"column:todoId"`
	PriorityId      *int   `json:"id"`
	UserName        string `json:"user_name"`
	Email           string `json:"email"`
	PhoneNo         *int   `json:"phone_no"`
	TodoTitle       string `json:"todo_title"`
	TodoDescription string `json:"todo_description"`
	TodoImage       string `json:"todo_image"`
	Priority        string `json:"priority"`
	Status          string `json:"status"`
}
