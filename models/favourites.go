package models

type Favourite struct {
    Base
	userId          string `json:"userId"` 
	todoId          string `json:"todoId"` 

}

func (m Favourite) TableName() string {
    return "favourites"
}

func (m Favourite) ToMap() map[string]interface{} {
    return map[string]interface{}{
        "userId":           m.userId,
		"todoId":			m.todoId,
}
}