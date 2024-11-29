package services

import (
	"Crud_fiber_Go/models"
	"Crud_fiber_Go/utils"
	"errors"
	"log"
)

func ValidateUser(user *models.User) error {
	if user.Name == "" || user.Email == "" {
		return errors.New("name and email are required")
	}
	return nil
}

func CreateUser(user *models.User) error {
	// Save user to database
	// DB logic here...
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Fatal(err)
	}

	user.Password = hashedPassword
	return models.DB.Create(&user).Error
}

func UpdateUser(id uint, user *models.User) error {
	user.ID = id
	return models.DB.Save(&user).Error
}

func GetUsers() ([]models.User, error) {
	var users []models.User
	// DB logic here...

	err := models.DB.Find(&users).Error
	return users, err
}

func DeleteUser(id string) error {
	err := models.DB.Where("id = ?", id).Delete(&models.User{}).Error
	return err
}
