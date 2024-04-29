package main

import (
	"github.com/ashrafatef/urlshortening/server"
)

func main() {
	server := server.SetupServer()
	server.Listen(":3000")
}
