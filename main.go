package main

import (
	"fmt"
	"os"

	"github.com/ashrafatef/urlshortening/server"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	server := server.SetupServer()
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	server.Listen(port)
}
