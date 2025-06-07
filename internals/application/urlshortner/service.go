package urlshortner

type Service struct {

}

func NewService() (*Service, error) {
	return &Service{}, nil
}

func (serivce Service) ShortenUrl(mainUrl string) (shortCode string, err error) {
	return "shortCode", nil
}

func (service Service) ExpandUrl(shortCode string) (mainUrl string, err error) {
	return "mainUrl", nil
}