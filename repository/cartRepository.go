package repository

import (
	"fullstack-capstone-backend/models"
	"fullstack-capstone-backend/config"
	"gorm.io/gorm"
	"strconv"
)

// GetCartByUserID fetches the cart items for a given user
func GetCartByUserID(userID string) ([]models.Cart, error) {
	var cartItems []models.Cart
	err := config.DB.Preload("Product").Where("user_id = ?", userID).Find(&cartItems).Error
	return cartItems, err
}

// ClearCartByUserID removes all cart items for a given user
func ClearCartByUserID(userID string) error {
	err := config.DB.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
	return err
}

// RemoveProductFromCart removes a specific product from the user's cart
func RemoveProductFromCart(userID string, productID string) error {
	err := config.DB.Where("user_id = ? AND product_id = ?", userID, productID).Delete(&models.Cart{}).Error
	return err
}

// AddProductToCart adds a product to the user's cart
func AddProductToCart(userID string, productID string, quantity int) error {
	var cartItem models.Cart

	// Convert userID and productID to uint
	userIDUint, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		return err
	}
	productIDUint, err := strconv.ParseUint(productID, 10, 32)
	if err != nil {
		return err
	}

	// Check if the product already exists in the user's cart
	err = config.DB.Where("user_id = ? AND product_id = ?", uint(userIDUint), uint(productIDUint)).First(&cartItem).Error
	if err != nil {
		// If the product doesn't exist, add a new cart item
		if err == gorm.ErrRecordNotFound {
			newCartItem := models.Cart{
				UserID:    uint(userIDUint),
				ProductID: uint(productIDUint),
				Quantity:  quantity,
			}
			return config.DB.Create(&newCartItem).Error
		}
		// Return any other error
		return err
	}

	// If the product already exists in the cart, update the quantity
	cartItem.Quantity += quantity
	return config.DB.Save(&cartItem).Error
}
