package repositories

import (
	"github.com/ashrafatef/urlshortening/application/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UrlCreateAttrs struct {
	Url       string
	HashedUrl string
}

type Urls struct {
	gorm.Model
	OriginalUrl string `json:"originalUrl"`
	ShortUrl    string `json:"shortUrl",gorm:"index"`
}

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) Create(input UrlCreateAttrs) (*Urls, error) {
	url := &Urls{
		OriginalUrl: input.Url,
		ShortUrl:    input.HashedUrl,
	}

	if err := r.db.Create(url).Error; err != nil {
		logrus.Error("Error creating url: ", err)
		return nil, errors.NewApplicationError("Error creating url")
	}

	return url, nil
}

func (r *UrlRepository) FindOne(id string) (*Urls, error) {
	var url *Urls

	result := r.db.Take(&url, "short_url = ?", id)
	if url == nil {
		return nil, errors.NewNotFoundError("Url not found")
	}

	if result.Error != nil {
		logrus.Error("Error finding url: ", result.Error)
		return nil, errors.NewApplicationError("Error finding url")
	}

	return url, nil
}
