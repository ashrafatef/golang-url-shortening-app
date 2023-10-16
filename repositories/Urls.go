package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type Urls struct {
	gorm.Model
	Url string `json:"url"`
}

type UrlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) *UrlRepository {
	return &UrlRepository{
		db: db,
	}
}

func (r *UrlRepository) Create(input Urls) uint {
	url := Urls{Url: input.Url}
	err := r.db.Create(&url)
	if err != nil {
		fmt.Println("Error creating url")
	}
	return url.ID
}

func (r *UrlRepository) Find() []Urls {
	var urls []Urls
	err := r.db.Find(&urls).Error
	if err != nil {
		fmt.Println("Error creating url")
		return []Urls{}
	}
	return urls
}

func (r *UrlRepository) FindOne(id uint) Urls {
	var url Urls
	err := r.db.Find(&url, "id = ?", id).Error
	if err != nil {
		fmt.Println("Error get url")
		return Urls{}
	}
	return url
}
