package domain


type ShortenedUrl struct {
	ID int64
	UserID int64
	ShortCode string
	MainUrl string
	ClickCount int64
	CreatedAt int64
}