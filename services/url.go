package services

import (
	"url-shortener/models"
	"url-shortener/repositories"

	"github.com/google/uuid"
)

type UrlService interface {
	Create(originalURL string) (string, error)

	Get(shortURL string) (models.Url, error)
}

func DefaultUrlService(repo repositories.UrlRepository) UrlService {
	return &urlService{
		repo,
	}
}

type urlService struct {
	repo repositories.UrlRepository
}

func (s *urlService) Create(originalURL string) (string, error) {
	urlUUID := uuid.Must(uuid.NewRandom())
	shortURL := urlUUID.String()[:8]

	url, _ := s.repo.GetUrl(shortURL)
	if url.ShortURL == "" {
		shortURL = urlUUID.String()[:8]
	}

	url = models.Url{}
	url.OriginalURL = originalURL
	url.ShortURL = shortURL

	err := s.repo.Create(url)
	if err != nil {
		return "", err
	}

	return shortURL, nil
}

func (s *urlService) Get(shortURL string) (models.Url, error) {
	url, err := s.repo.GetUrl(shortURL)
	if err != nil {
		return models.Url{}, err
	}

	return url, nil
}
