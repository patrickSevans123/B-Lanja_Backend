package repository

import (
	"errors"
	"fullstack-capstone-backend/models"
	"fullstack-capstone-backend/config"
)

// GetAllProducts retrieves all products
func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}

// GetProductByID retrieves a product by its ID
func GetProductByID(id string) (models.Product, error) {
	var product models.Product
	err := config.DB.First(&product, id).Error
	return product, err
}

// CreateProduct creates a new product in the database
func CreateProduct(product *models.Product) error {
	err := config.DB.Create(product).Error
	return err
}

// UpdateProduct updates a product by its ID only if the user owns the product
func UpdateProduct(id string, product *models.Product, userID uint) error {
	var existingProduct models.Product
	err := config.DB.First(&existingProduct, id).Error
	if err != nil {
		return err
	}

	// Check if the current user is the owner of the product
	if existingProduct.UserID != userID {
		return errors.New("you are not the owner of this product")
	}

	// Update existing product details
	existingProduct.Name = product.Name
	existingProduct.Price = product.Price
	existingProduct.Stock = product.Stock

	return config.DB.Save(&existingProduct).Error
}

// DeleteProduct deletes a product by its ID only if the user owns the product
func DeleteProduct(id string, userID uint) error {
	var product models.Product
	err := config.DB.First(&product, id).Error
	if err != nil {
		return err
	}

	// Check if the current user is the owner of the product
	if product.UserID != userID {
		return errors.New("you are not the owner of this product")
	}

	// Delete the product
	return config.DB.Delete(&product).Error
}
