package repository

import (
	"fullstack-capstone-backend/models"
	"fullstack-capstone-backend/config"
)

// CreateTransaction creates a new transaction in the database
func CreateTransaction(transaction *models.Transaction) error {
	return config.DB.Create(transaction).Error
}

// GetTransactionsByUserID retrieves all transactions for a specific user
func GetTransactionsByUserID(userID string) ([]models.Transaction, error) {
	var transactions []models.Transaction

	// Query to find all transactions for the user ID
	if err := config.DB.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}

	return transactions, nil
}
