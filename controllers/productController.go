package controllers

import (
	"net/http"
	"fullstack-capstone-backend/models"
	"fullstack-capstone-backend/repository"
	"github.com/gin-gonic/gin"
)

// GetAllProducts retrieves all products (items)
func GetAllProducts(c *gin.Context) {
	products, err := repository.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProductByID retrieves a product by its ID
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	product, err := repository.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

// CreateProduct creates a new product and associates it with the creator (user_id)
func CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Save the product in the database
    if err := repository.CreateProduct(&product); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusCreated, product)
}



// UpdateProduct allows a user to update a product if they are the owner
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := repository.UpdateProduct(id, &product, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the owner of this product or product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}

// DeleteProduct allows a user to delete a product if they are the owner
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetUint("user_id")

	err := repository.DeleteProduct(id, userID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the owner of this product or product not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
