package repository

import (
	"fullstack-capstone-backend/models"
	"fullstack-capstone-backend/config"
)

// GetAllUsers retrieves all users from the database
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := config.DB.Find(&users).Error
	return users, err
}

// GetUserByID retrieves a user by their ID
func GetUserByID(id string) (models.User, error) {
	var user models.User
	err := config.DB.First(&user, id).Error
	return user, err
}

func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// GetUserByEmail retrieves a user by their email
func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}