package models

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	OriginalURL string `json:"originalUrl" gorm:"original_url"`
	ShortURL    string `json:"shortUrl" gorm:"short_url"`
}
