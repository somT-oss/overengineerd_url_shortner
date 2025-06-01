package adapters

import (
	"log"

	"github.com/somT-oss/urlshortner/internals/application/core/domain"
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


func (adapter Adapter) GetUserById(id int64) (domain.User, error) {
	var userEntity User
	res := adapter.db.First(&userEntity, id)
	user := domain.User{
		ID: int64(userEntity.ID),
		Name: userEntity.Name,
		Email: userEntity.Email,
		Password: userEntity.Password,
		CreatedAt: int(userEntity.CreatedAt.UnixNano()),
		UpdatedAt: int(userEntity.CreatedAt.UnixNano()),
	}

	return user, res.Error
}

func (adapter Adapter) GetUserByEmail(email string) (domain.User, error) {
	var userEntity User
	res := adapter.db.First(&userEntity, email)
	user := domain.User{
		ID: int64(userEntity.ID),
		Name: userEntity.Name,
		Email: userEntity.Email,
		Password: userEntity.Password,
		CreatedAt: int(userEntity.CreatedAt.UnixNano()),
		UpdatedAt: int(userEntity.CreatedAt.UnixNano()),
	}

	return user, res.Error	
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


func (adapter Adapter) SaveUser(user *domain.User) error {
	userModel := User{
		Name: user.Name,
		Email: user.Email,
		Password: user.Password,
	}
	res := adapter.db.Create(&userModel)	
	if res.Error == nil {
		user.ID = int64(userModel.ID)
	}
	return res.Error
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