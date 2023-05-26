package repository

import (
	"api-test/interfaces"
	"api-test/models"
	"fmt"

	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &repository{
		DB: db,
	}
}

func (r *repository) LoginRepository(username string) (user models.User, err error) {
	err = r.DB.Raw("SELECT * FROM users WHERE username = ?", username).Scan(&user).Error

	if err != nil {
		fmt.Println("this this the error : ", err)
	}

	return user, nil
}

func (r *repository) RegisterRepository(user models.User) error {
	response := r.DB.Create(&user)

	if response.Error != nil {
		return response.Error
	}

	return nil
}
