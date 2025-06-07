package api

import (
	"log"

	"github.com/somT-oss/urlshortner/internals/application/domain"
	"github.com/somT-oss/urlshortner/internals/ports"
)

type Application struct {
	db ports.DBPort
	service UrlShortnerService
}


func NewApplication(db ports.DBPort, service UrlShortnerService) *Application {
	return &Application{
		db: db,
		service: service,
	}
}


func (app Application) GenerateShortUrl(mainUrl string) (domain.ShortenedUrl, error) {
	if len(mainUrl) == 0 {
		log.Fatalf("the length of the string is less than 0")
	}
	shortUrl, err := app.service.ShortenUrl(mainUrl)
	if err != nil {
		log.Fatalf("the error %v occurred", err)
	}
	shortenedUrl := domain.ShortenedUrl{
		ID: 0,
		UserID: 0,
		ShortCode: shortUrl,
		MainUrl: mainUrl,
		ClickCount: 0,
		CreatedAt: 0, 
	}
	err = app.db.SaveShortenedUrl(&shortenedUrl)
	if err != nil {
		log.Fatalf("the error %v occurred", err)
	}
	return shortenedUrl, nil
}
