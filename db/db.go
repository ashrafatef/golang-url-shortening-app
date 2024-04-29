package db

import (
	"fmt"

	"github.com/ashrafatef/urlshortening/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Connect() *gorm.DB {

	if connection != nil {
		return connection
	}

	dns := "host=localhost user=postgres dbname=url_shortening port=5432 sslmode=disable"

	connection, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error creating database")
	}

	connection.Migrator().AutoMigrate(&repositories.Urls{})

	return connection
}
