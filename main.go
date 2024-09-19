package main

import (
    "github.com/gin-gonic/gin"
    "fullstack-capstone-backend/config"
    "fullstack-capstone-backend/routes"
)

func main() {
    // Initialize database
    config.Initialize()

    // Create a new Gin router
    router := gin.Default()

    // Initialize routes with the router and database connection
    routes.InitRoutes(router, config.DB)

    // Start the server
    router.Run(":8080")
}
