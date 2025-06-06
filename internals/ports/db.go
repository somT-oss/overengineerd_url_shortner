package ports

import "github.com/somT-oss/urlshortner/internals/application/core/domain"


type DBPort interface {
	GetUrlById(id int) (domain.ShortenedUrl, error)
	GetUrlByShortUrl(shorUrl string) (domain.ShortenedUrl, error)
	GetUserById(id int) (domain.User, error)
	GetUserByEmail(email string) (domain.User, error)
	SaveUser()
	SaveShortenedUrl()
}