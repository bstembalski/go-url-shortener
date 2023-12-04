package responses

import "time"

type CreateUrl struct {
	OriginalURL string `json:"originalURL"`
}

type GetUrl struct {
	OriginalURL string `json:"originalUrl"`
	ShortURL    string `json:"shortUrl"`
}

type GetAllUrls struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	OriginalURL string `json:"originalUrl" gorm:"original_url"`
	ShortURL    string `json:"shortUrl" gorm:"short_url"`
}

type GetUrlRequest struct {
	ShortURL string `json:"shortUrl"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
