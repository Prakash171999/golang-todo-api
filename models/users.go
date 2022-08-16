package models

type User struct {
	Base
	Email       string `json:"email" binding:"required,email" gorm:"unique"`
	PhoneNumber int    `json:"phone_number"`
	FullName    string `json:"full_name"`
	Password    []byte `json:"-"`
}

// TableName gives table name of model
func (m User) TableName() string {
	return "users"
}

// ToMap convert User to map
func (m User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":        m.Email,
		"phone_number": m.PhoneNumber,
		"full_name":    m.FullName,
		"password":     m.Password,
	}
}
