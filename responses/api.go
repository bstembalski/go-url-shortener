package responses

type CreateUrl struct {
	OriginalURL string `json:"originalURL"`
}

type GetUrl struct {
	OriginalURL string `json:"originalUrl"`
	ShortURL    string `json:"shortUrl"`
}

type GetUrlRequest struct {
	ShortURL string `json:"shortUrl"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
