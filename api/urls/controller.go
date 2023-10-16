package urls

import (
	"strconv"

	"github.com/ashrafatef/urlshortening/repositories"
	"github.com/gofiber/fiber/v2"
)

type UrlController struct {
	urlRepo *repositories.UrlRepository
}

func NewUrlController(urlRepo *repositories.UrlRepository) *UrlController {
	return &UrlController{
		urlRepo: urlRepo,
	}
}

func (u *UrlController) List(c *fiber.Ctx) error {

	urls := u.urlRepo.Find()
	if len(urls) == 0 {
		println("Could not find urls")
	}
	return c.JSON(urls)
}

func (u *UrlController) Create(c *fiber.Ctx) error {
	input := new(repositories.Urls)

	if err := c.BodyParser(input); err != nil {
		return err
	}
	id := u.urlRepo.Create(*input)

	return c.JSON(id)
}

func (u *UrlController) Get(c *fiber.Ctx) error {
	param := c.Params("id")
	id, _ := strconv.ParseUint(param, 10, 10)

	url := u.urlRepo.FindOne(uint(id))

	return c.JSON(url)
}
