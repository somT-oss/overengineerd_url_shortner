package ports

import (
	"github.com/somT-oss/urlshortner/internals/application/domain"
)

type APIPort interface {
	GenerateShortUrl(mainUrl string) (domain.ShortenedUrl, error)
}