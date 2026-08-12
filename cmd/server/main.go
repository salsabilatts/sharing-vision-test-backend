package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/salsabilatts/sharing-vision-test/config"
	"github.com/salsabilatts/sharing-vision-test/internal/handler"
	"github.com/salsabilatts/sharing-vision-test/internal/repository"
	"github.com/salsabilatts/sharing-vision-test/internal/service"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatal("Failed to load environment variables:", err)
	}

	db, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get SQL database:", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connected successfully")

	// Dependencies
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postHandler := handler.NewPostHandler(postService)

	// Router
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// Health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Article routes
	router.POST("/article/", postHandler.Create)
	router.GET("/article/:param/:offset", postHandler.GetAll)
	router.GET("/article/:param", postHandler.GetByID)
	router.PUT("/article/:id", postHandler.Update)
	router.PATCH("/article/:id", postHandler.Update)
	router.DELETE("/article/:id", postHandler.Delete)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running on port %s", port)

	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
