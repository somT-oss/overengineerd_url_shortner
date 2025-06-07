package adapters

import (
	"log"

	"github.com/somT-oss/urlshortner/internals/application/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ShortenedUrl struct {
	gorm.Model
	UserID int64
	MainUrl string
	ShortCode string
	ClickCount int64
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(databaseUrl string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database. Error %v occurred", err)
		return nil, err
	}
	err = db.AutoMigrate(&ShortenedUrl{})
	if err != nil {
		log.Fatalf("failed to migrate the schema to the db. Error %v occurred", err)
	}
	return &Adapter{
		db: db,
	}, nil
}

func (adapter Adapter) GetUrlById(id int64) (domain.ShortenedUrl, error) {
	var shortenedUrlEntity ShortenedUrl
	res := adapter.db.First(&shortenedUrlEntity, id) 
	shortenedUrl := domain.ShortenedUrl{
		ID: int64(shortenedUrlEntity.ID),
		UserID: shortenedUrlEntity.UserID,
		ClickCount: shortenedUrlEntity.ClickCount,
		CreatedAt: shortenedUrlEntity.CreatedAt.UnixNano(),
	}
	return shortenedUrl, res.Error
}


func (adapter Adapter) GetUrlByShortCode(shortCode string) (domain.ShortenedUrl, error) {
	var shortenedUrlEntity ShortenedUrl
	res := adapter.db.First(&shortenedUrlEntity, shortCode)
	shortenedUrl := domain.ShortenedUrl{
		ID: shortenedUrlEntity.UserID,
		UserID: shortenedUrlEntity.UserID,
		ClickCount: shortenedUrlEntity.ClickCount,
		CreatedAt: shortenedUrlEntity.CreatedAt.UnixNano(),
	}
	return shortenedUrl, res.Error
}

func (adapter Adapter) SaveShortenedUrl(shortenedUrl *domain.ShortenedUrl) error {
	shortenedUrlModel := ShortenedUrl{
		UserID: shortenedUrl.UserID,
		MainUrl: shortenedUrl.MainUrl,
		ClickCount: shortenedUrl.ClickCount,
	}
	res := adapter.db.Create(&shortenedUrlModel)
	if res.Error == nil {
		shortenedUrl.ID = int64(shortenedUrl.ID)
	}
	return res.Error
}