package main

import (
	"fmt"

	"log"

	"github.com/joho/godotenv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"light_novel/backend/controllers"
	"light_novel/backend/database"
	"light_novel/backend/middlewares"
	"light_novel/backend/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("File .env tidak ditemukan, menggunakan environment variable dari sistem")
	}

	database.ConnectPostgreDB()
	database.DB.AutoMigrate(&models.User{})
	fmt.Println("Database connected and migrated")

	r := gin.Default()
	r.Use(cors.Default())

	protected := r.Group("/api")
	protected.Use(middlewares.AuthorizeUser())
	// protected.GET("/profile", controllers.Profile)

	r.POST("/login", controllers.Login)
	r.Run()
}
