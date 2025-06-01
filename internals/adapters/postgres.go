package adapters

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
}

type ShortenedUrl struct {
	gorm.Model
	UserID int64
	MainUrl string
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
	err = db.AutoMigrate(&User{}, &ShortenedUrl{})
	if err != nil {
		log.Fatalf("failed to migrate the schema to the db. Error %v occurred", err)
	}
	return &Adapter{
		db: db,
	}, nil
}