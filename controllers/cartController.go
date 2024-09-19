package controllers

import (
	"net/http"
	"fullstack-capstone-backend/repository"
	"github.com/gin-gonic/gin"
)

// GetCartByUserID retrieves all cart items for a user
func GetCartByUserID(c *gin.Context) {
	userID := c.Param("user_id")
	cart, err := repository.GetCartByUserID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cart)
}

// RemoveProductFromCart removes a specific product from the cart
func RemoveProductFromCart(c *gin.Context) {
	userID := c.Param("user_id")
	productID := c.Param("product_id")

	err := repository.RemoveProductFromCart(userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product removed from cart"})
}

// AddProductToCart adds a product to the user's cart
func AddProductToCart(c *gin.Context) {
	var requestData struct {
		UserID    string `json:"user_id" binding:"required"`
		ProductID string `json:"product_id" binding:"required"`
		Quantity  int    `json:"quantity" binding:"required,min=1"`
	}

	// Parse JSON request data
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add product to the user's cart
	err := repository.AddProductToCart(requestData.UserID, requestData.ProductID, requestData.Quantity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add product to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product added to cart successfully"})
}

