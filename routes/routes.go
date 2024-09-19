package routes

import (
	"fullstack-capstone-backend/controllers"
	"fullstack-capstone-backend/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"gorm.io/gorm"
)

// InitRoutes initializes all the routes
func InitRoutes(r *gin.Engine, db *gorm.DB) {
	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"}, // Replace with the origin of your frontend
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// Public routes
	r.GET("/products", controllers.GetAllProducts)
	r.GET("/products/:id", controllers.GetProductByID)

	// User-related routes
	r.GET("/users", controllers.GetAllUsers)       // Get all users
	r.GET("/users/:id", controllers.GetUserByID)   // Get a user by ID
	r.POST("/register", controllers.RegisterUser)  // Register a new user
	r.POST("/login", controllers.LoginUser)        // Login a user

	// Apply authentication middleware to the following routes
	r.Use(middleware.AuthMiddleware())
	r.POST("/products", controllers.CreateProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	r.GET("/carts/:user_id", controllers.GetCartByUserID)
	r.POST("/checkout", controllers.Checkout)
	r.DELETE("/carts/:product_id", controllers.RemoveProductFromCart)
	r.POST("/carts", controllers.AddProductToCart)

	// Transaction-related routes
	r.GET("/transactions/:user_id", controllers.GetTransactionsByUserID)  // Get transactions by user ID
}
