package controllers

import (
	"net/http"
	"strconv"
	"fullstack-capstone-backend/models"
	"fullstack-capstone-backend/repository"

	"github.com/gin-gonic/gin"
)

// Checkout processes the checkout and clears the cart
func Checkout(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
        return
    }

    userIDStr := strconv.Itoa(int(userID.(float64)))

    cartItems, err := repository.GetCartByUserID(userIDStr)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cart"})
        return
    }

    if len(cartItems) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Cart is empty"})
        return
    }

    var totalAmount float64
    for _, item := range cartItems {
        // Fetch product details
        var product models.Product
        if err := repository.GetProductByID(strconv.Itoa(int(item.ProductID)), &product); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
            return
        }

        // Calculate total amount
        totalAmount += float64(item.Quantity) * product.Price
    }

    transaction := models.Transaction{
        UserID: uint(userID.(float64)),
        Amount: totalAmount,
    }

    if err := repository.CreateTransaction(&transaction); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
        return
    }

    if err := repository.ClearCartByUserID(userIDStr); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear cart"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message":    "Checkout successful",
        "transaction": transaction,
    })
}

// GetTransactionsByUserID retrieves all transactions for a specific user by their user ID
func GetTransactionsByUserID(c *gin.Context) {
	// Get user_id from the URL parameter
	userID := c.Param("user_id")

	// Fetch transactions by user ID from the repository
	transactions, err := repository.GetTransactionsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the transactions
	c.JSON(http.StatusOK, transactions)
}