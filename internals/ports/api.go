package ports

import (
	"github.com/somT-oss/urlshortner/internals/application/core/domain"
)

type APIPort interface {
	GenerateShortUrl(mainUrl string) (domain.ShortenedUrl, error)
	CreateUser(name, email, password string) (domain.User, error)
}