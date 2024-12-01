package main

import (
	"github.com/ashrafatef/urlshortening/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	server := server.SetupServer()
	server.Listen(":3000")
}
