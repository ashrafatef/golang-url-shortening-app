package db

import (
	"fmt"

	"github.com/ashrafatef/urlshortening/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	dns := "host=localhost user=postgres dbname=url_shortening port=5432 sslmode=disable"

	Conn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error creating database")
	}
	
	Conn.Migrator().AutoMigrate(&repositories.Urls{})

	return Conn
}
