package repositories

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"gorm.io/gorm"
)

type UrlInput struct {
	Url string
}

type Urls struct {
	gorm.Model
	OriginalUrl string `json:"originalUrl"`
	ShortUrl    string `json:"shortUrl"`
}

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) Create(input UrlInput) Urls {
	hasher := md5.New()
	hasher.Write([]byte(input.Url))
	hashedUrl := hex.EncodeToString(hasher.Sum(nil))
	url := Urls{
		OriginalUrl: input.Url,
		ShortUrl:    hashedUrl,
	}
	err := r.db.Create(&url)
	if err != nil {
		fmt.Println("Error creating url")
	}
	return url
}

func (r *UrlRepository) FindOne(id string) Urls {
	var url Urls
	err := r.db.Find(&url, "short_url = ?", id).Error
	if err != nil {
		fmt.Println("Error get url", err)
		return Urls{}
	}
	return url
}
