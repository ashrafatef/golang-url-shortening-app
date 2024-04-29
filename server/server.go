package server

import (
	"net/http"

	"github.com/ashrafatef/urlshortening/api/urls"
	"github.com/ashrafatef/urlshortening/db"
	"github.com/gofiber/fiber/v2"
)

func SetupServer() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		StrictRouting:     true,
	})
	db.Connect()

	app = urls.SetUpUrlRoutes(app)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	return app
}
