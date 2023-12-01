package api

import (
	"log"
	"net/http"
	"url-shortener/responses"
	"url-shortener/services"

	"github.com/gin-gonic/gin"
)

type UrlController interface {
	Create(c *gin.Context)
	Get(c *gin.Context)
}

func DefaultUrlontroller(urlService services.UrlService) UrlController {
	return &urlController{
		urlService,
	}
}

type urlController struct {
	urlService services.UrlService
}

func (ctr *urlController) Create(c *gin.Context) {
	resUrl := responses.CreateUrl{}

	err := c.BindJSON(&resUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	shortUrl, err := ctr.urlService.Create(resUrl.OriginalURL)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"shortUrl": shortUrl})
}

func (ctr *urlController) Get(c *gin.Context) {
	var request responses.GetUrlRequest

	if err := c.BindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	shortUrl := request.ShortURL

	urlInfo, err := ctr.urlService.Get(shortUrl)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	info := responses.GetUrl{}
	info.OriginalURL = urlInfo.OriginalURL
	info.ShortURL = urlInfo.ShortURL

	c.JSON(http.StatusOK, gin.H{"URL": info})
}
