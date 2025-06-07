package ports

import "github.com/somT-oss/urlshortner/internals/application/domain"


type DBPort interface {
	GetUrlById(id int) (domain.ShortenedUrl, error)
	GetUrlByShortCode(shortCode string) (domain.ShortenedUrl, error)
	SaveShortenedUrl(*domain.ShortenedUrl) error
}