package urls

import (
	"github.com/ashrafatef/urlshortening/repositories"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetUpUrlRoutes(app *fiber.App, conn *gorm.DB) {

	urlRepo := repositories.NewUrlRepository(conn)
	urlController := NewUrlController(urlRepo)

	app.Post("/", urlController.Create)
	app.Get("/", urlController.List)
	app.Get("/:id", urlController.Get)
}
