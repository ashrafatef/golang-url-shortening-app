package urls

import (
	"fmt"
	"net/http"

	"github.com/ashrafatef/urlshortening/repositories"
	"github.com/gofiber/fiber/v2"
)

type UrlInput struct {
	OriginalUrl string
}

type UrlController struct {
	urlRepo *repositories.UrlRepository
}

func NewUrlController(urlRepo *repositories.UrlRepository) *UrlController {
	return &UrlController{
		urlRepo: urlRepo,
	}
}

func (u *UrlController) Create(c *fiber.Ctx) error {
	input := new(UrlInput)
	if err := c.BodyParser(input); err != nil {
		return err
	}
	fmt.Printf("%+v", input)

	id := u.urlRepo.Create(repositories.UrlInput{
		Url: input.OriginalUrl,
	})

	return c.JSON(id)
}

func (u *UrlController) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	fmt.Printf("%+v\n", id)
	url := u.urlRepo.FindOne(id)
	fmt.Printf("yrl is %+v\n", url)

	return c.Redirect(url.OriginalUrl, http.StatusMovedPermanently)
}
