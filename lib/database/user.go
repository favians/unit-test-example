package database

import (
	"mvc/config"
	"mvc/middleware"
	"mvc/models"
)

func GetSingleUser(id int) (interface{}, error) {
	var users *models.Users

	if e := config.DB.First(&users, id).Error; e != nil {
		return nil, e
	}

	return users, nil
}

func GetUsers() (interface{}, error) {
	var users []models.Users

	if e := config.DB.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func InsertUser(user models.Users) error {
	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(user models.Users) error {
	if err := config.DB.Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func LoginUsers(user *models.Users) (interface{}, error) {

	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middleware.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
