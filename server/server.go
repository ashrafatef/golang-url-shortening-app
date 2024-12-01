package server

import (
	"fmt"
	"net/http"

	"github.com/ashrafatef/urlshortening/api/urls"
	"github.com/ashrafatef/urlshortening/db"
	"github.com/ashrafatef/urlshortening/errors"
	"github.com/gofiber/fiber/v2"
)

func SetupServer() *fiber.App {
	app := fiber.New(fiber.Config{
		EnablePrintRoutes: true,
		StrictRouting:     true,
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			if validationErr, ok := err.(*errors.ValidationError); ok {
				fmt.Println("errror", err)

				return c.Status(validationErr.StatusCode).JSON(fiber.Map{
					"message": validationErr.Message,
					"fields":  validationErr.Fields,
				})
			}
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})
	db.Connect()

	app = urls.SetUpUrlRoutes(app)
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	return app
}
