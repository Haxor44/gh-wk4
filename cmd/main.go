package main

import (
	"E-matBackend/internal/handlers"
	"E-matBackend/internal/repositories/mysql"
	"E-matBackend/internal/repositories/redis"
	"E-matBackend/internal/services"
	"E-matBackend/pkg/database"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// initialize db
	db := database.MySQLConnection()
	defer db.Close()

	rdb := database.Redisconnection()
	defer rdb.Close()

	// initialize repositories
	productRepo := mysql.NewProductRepository(db)
	cacheRepo := redis.NewCacheRepository(rdb)

	// initialize service
	productService := services.NewProductService(productRepo, cacheRepo)

	product, err := productService.GetProduct(1)
	if err != nil {
		log.Fatalf("Error fetching product: %v", err)
	}

	fmt.Printf("Product: %+v\n", product)

	// Initialize handler
	productHandler := handlers.NewProductHandler(*productService)

	// Create Gin router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins (adjust for production)
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Define routes
	router.GET("/products/:id", productHandler.GetProduct)
	router.GET("/products", productHandler.GetProducts)

	// Start server
	log.Println("Server running on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
