package repository

import (
	"boilerplate-api/infrastructure"
	"boilerplate-api/models"
)

type UserAuthRepository struct {
	db     infrastructure.Database
	logger infrastructure.Logger
}

func NewUserAuthRepository(db infrastructure.Database, logger infrastructure.Logger) UserAuthRepository {
	return UserAuthRepository{
		db:     db,
		logger: logger,
	}
}

func (c UserAuthRepository) Register(User models.User) (models.User, error) {

	// var existingUser models.User

	// emailExists := c.db.DB.Where("email = ?", User.Email).First(&existingUser)

	// if emailExists.RowsAffected > 0 {
	// 	fmt.Println("Email exists. Please use another email address.")
	// 	return existingUser, c.db.DB.Error
	// }

	return User, c.db.DB.Create(&User).Error

}

func (c *UserAuthRepository) GetUserFromEmail(userEmail string) (loggedInUser models.User, error error) {
	return loggedInUser, c.db.DB.Where("email = ?", userEmail).First(&loggedInUser).Error
}

// Update/Reset user password
func (c *UserAuthRepository) ResetPassword(User models.User, password string) error {
	return c.db.DB.Model(&models.User{}).
		Where("email = ?", User.Email).
		Updates(map[string]interface{}{
			"password": password,
		}).Find(&User).Error
}

// func (c UserAuthRepository) Login(user models.User) bool {

// password := bcrypt.CompareHashAndPassword([]byte(user.Password), user.Password)
// password, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
// decryotPwd := string(password)

// if err := c.db.DB.Model(&models.User{}).Where("email = ?", user.Email).Where("password = ?", decryotPwd).First(&user).Error; err != nil {
// 	return false
// }
// return true
// }
