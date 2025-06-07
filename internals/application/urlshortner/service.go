package urlshortner

type Service struct {

}

func NewService() *Service {
	return &Service{}
}

func (serivce Service) ShortenUrl(mainUrl string) (shortCode string, err error) {
	return "shortCode", nil
}

func (service Service) ExpandUrl(shortCode string) (mainUrl string, err error) {
	return "mainUrl", nil
}