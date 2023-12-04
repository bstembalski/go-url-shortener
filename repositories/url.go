package repositories

import (
	"log"
	"url-shortener/models"

	"gorm.io/gorm"
)

type UrlRepository interface {
	Create(url models.Url) error

	GetUrl(shortUrl string) (models.Url, error)
	GetAllUrls() (models.Url, error)
}

func DefaultUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepo{
		db,
	}
}

type urlRepo struct {
	db *gorm.DB
}

func (r *urlRepo) Create(url models.Url) error {
	result := r.db.Create(&url)

	if result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}

	return nil
}

func (r *urlRepo) GetUrl(shortUrl string) (models.Url, error) {
	url := models.Url{}

	result := r.db.First(&url, "short_url = ?", shortUrl)
	if result.Error != nil {
		return models.Url{}, result.Error
	}

	return url, nil
}

func (r *urlRepo) GetAllUrls() (models.Url, error) {
	var url models.Url

	result := r.db.Find(&url)
	if result.Error != nil {
		return models.Url{}, result.Error
	}

	return url, nil
}
