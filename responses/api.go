package responses

type CreateUrl struct {
	OriginalURL string `json:"originalUrl"`
}

type GetUrl struct {
	OriginalURL string `json:"originalUrl"`
	ShortURL    string `json:"shortUrl"`
}

type GetUrlRequest struct {
	ShortURL string `json:"shortUrl"`
}
