package main

import (
	"log"
	"url-shortener/api"
	"url-shortener/helpers"
	"url-shortener/models"
	"url-shortener/repositories"
	"url-shortener/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var postgresdsn = helpers.GetEnv("POSTGRES_DSN", "")

func main() {
	db, err := gorm.Open(postgres.Open(postgresdsn))
	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	db.AutoMigrate(&models.Url{})

	gin := gin.Default()
	// gin.Use(helmet.Default())

	urlRepo := repositories.DefaultUrlRepository(db)
	urlService := services.DefaultUrlService(urlRepo)
	controller := api.DefaultUrlontroller(urlService)

	urlGroup := gin.Group("/")
	{
		urlGroup.POST("/create", controller.Create)
		urlGroup.GET("/get", controller.Get)
	}

	err = gin.Run()
	if err != nil {
		log.Fatal(err)
	}
}
