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
	GetAllUrls(c *gin.Context)
}

func DefaultUrlController(urlService services.UrlService) UrlController {
	return &urlController{
		urlService,
	}
}

type urlController struct {
	urlService services.UrlService
}

// Create short URL
// @Summary Creates a short URL
// @ID create-url
// @Accept json
// @Produce json
// @Param body body responses.CreateUrl true "Create URL request body"
// @Success 201 {string} string "Short URL successfully created"
// @Failure 400,401,500 {object} responses.ErrorResponse "Error response"
// @Router /create [post]
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

// Get URL info
// @Summary Shows URL information
// @ID get-url
// @Accept json
// @Produce json
// @Param body body responses.GetUrlRequest true "Get URL request body"
// @Success 200 {object} responses.GetUrl "URL information"
// @Failure 400,401,500 {object} responses.ErrorResponse "Error response"
// @Router /get [get]
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

	c.JSON(http.StatusOK, info)
}

// Get all URLs
// @Summary Shows all urls
// @ID get-all-urls
// @Accept json
// @Produce json
// @Success 200 {array} responses.GetAllUrls "All urls"
// @Failure 400,401,500 {object} responses.ErrorResponse "Error response"
// @Router / [get]
func (ctr *urlController) GetAllUrls(c *gin.Context) {
	urls, err := ctr.urlService.GetAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusOK, urls)
}
