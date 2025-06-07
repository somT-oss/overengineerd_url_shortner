package api

type UrlShortnerService interface {
	ShortenUrl(mainUrl string) (shortCode string, err error)
	ExpandUrl(shortCode string) (mainUrl string, err error)
}