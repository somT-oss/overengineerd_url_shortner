package domain

import "time"

type ShortenedUrl struct {
	ID int64 `json:"id"`
	UserID int64 `json:"user_id"`
	MainUrl string `json:"main_url"`
	CreatedAt int64 `json:"created_at"`
	ClickCount int64 `json:"click_count"`
}

func NewShortenedUrl(mainUrl string) ShortenedUrl {
	return ShortenedUrl{
		CreatedAt: time.Now().Unix(),
		MainUrl: mainUrl,
		ClickCount: 0,
	}
}


