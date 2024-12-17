package db

import (
	"fmt"
	"os"

	"github.com/ashrafatef/urlshortening/infra/repositories"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func Connect() *gorm.DB {

	if connection != nil {
		return connection
	}

	host := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUserName := os.Getenv("DB_USER_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dns := fmt.Sprintf("host=%s user=%s dbname=%s port=%s sslmode=disable password=%s", host, dbUserName, dbName, dbPort, dbPassword)

	connection, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error While Opening Database Connection")
	}

	connection.Migrator().AutoMigrate(&repositories.Urls{})

	return connection
}
