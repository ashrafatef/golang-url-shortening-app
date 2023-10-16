package main

import (
	"github.com/ashrafatef/urlshortening/api/urls"
	"github.com/ashrafatef/urlshortening/db"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	conn := db.Connect()

	urls.SetUpUrlRoutes(app, conn)

	app.Listen(":3000")
}
