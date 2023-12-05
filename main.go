package main

import (
	"log"
	"url-shortener/api"
	"url-shortener/helpers"
	"url-shortener/models"
	"url-shortener/repositories"
	"url-shortener/services"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "url-shortener/docs"
)

// gin-swagger middleware

var postgresdsn = helpers.GetEnv("MYSQL_DSN", "root:root@tcp(127.0.0.1:3306)/url-shortener?charset=utf8mb4&parseTime=True&loc=Local")

// @title           Url shortener
// @version         1.0
// @description     This API allows you to shorten your links and redirect your webpage to a specified link.
// @host      localhost:8080
// @BasePath  /
func main() {
	db, err := gorm.Open(mysql.Open(postgresdsn))
	if err != nil {
		log.Fatal("Failed to connect DB")
	}
	db.AutoMigrate(&models.Url{})

	gin := gin.Default()

	gin.Use(helmet.Default())
	gin.Use(cors.Default())

	urlRepo := repositories.DefaultUrlRepository(db)
	urlService := services.DefaultUrlService(urlRepo)
	controller := api.DefaultUrlController(urlService)

	urlGroup := gin.Group("/")
	{
		urlGroup.GET("/", controller.GetAllUrls)
		urlGroup.GET("/get", controller.Get)
		urlGroup.POST("/create", controller.Create)
	}

	gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	err = gin.Run()
	if err != nil {
		log.Fatal(err)
	}
}
