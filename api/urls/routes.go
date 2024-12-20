package urls

import (
	"github.com/ashrafatef/urlshortening/infra/db"
	"github.com/ashrafatef/urlshortening/infra/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetUpUrlRoutes(app *fiber.App) *fiber.App {
	dbConn := db.Connect()
	urlRepo := repositories.NewUrlRepository(dbConn)
	urlController := NewUrlController(urlRepo)

	app.Post("/urls", urlController.Create)
	app.Get("/urls/:id", urlController.GetUrl)

	return app
}
