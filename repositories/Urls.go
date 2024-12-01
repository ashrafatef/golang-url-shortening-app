package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type UrlInput struct {
	Url       string
	HashedUrl string
}

type Urls struct {
	gorm.Model
	OriginalUrl string `json:"originalUrl",gorm:"index"`
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

func (r *UrlRepository) Create(input UrlInput) (*Urls, error) {
	// Validate input
	if input.Url == "" || input.HashedUrl == "" {
		return nil, fmt.Errorf("url and hashed url are required")
	}

	url := &Urls{
		OriginalUrl: input.Url,
		ShortUrl:    input.HashedUrl,
	}

	// Create record in database
	if err := r.db.Create(url).Error; err != nil {
		return nil, fmt.Errorf("error creating url: %w", err)
	}

	return url, nil
}

func (r *UrlRepository) FindOne(id string) (*Urls, error) {
	var url Urls
	err := r.db.First(&url, "short_url = ?", id).Error
	if err != nil {
		fmt.Println("Error get url", err)
		return nil, err
	}
	return &url, nil
}
