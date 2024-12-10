package urls

import (
	"crypto/rand"
	"math/big"
	"net/http"

	"github.com/ashrafatef/urlshortening/errors"
	"github.com/ashrafatef/urlshortening/repositories"
	"github.com/ashrafatef/urlshortening/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sirupsen/logrus"
)

type UrlInput struct {
	OriginalUrl string `validate:"required,url" json:"original_url"`
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

	if err := c.BodyParser(&input); err != nil {
		log.Error(err)
		return err
	}

	if errs := validations.Validation(input); len(errs) != 0 {
		logrus.Error("Error validating input: ", errs)
		return errors.NewValidationError(errs)
	}
	var shortUrl string
	for ok := true; ok; {
		hashedUrl, _ := generateRandomCode(8)
		url, _ := u.urlRepo.FindOne(hashedUrl)

		if url != nil {
			continue
		}
		_, error := u.urlRepo.Create(repositories.UrlCreateAttrs{
			Url:       input.OriginalUrl,
			HashedUrl: hashedUrl,
		})

		if error != nil {
			logrus.Error("Error creating url: ", error)
			return c.Status(http.StatusInternalServerError).JSON("Internal Server Error")
		}
		shortUrl = hashedUrl
		ok = false
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"shortUrl": shortUrl,
	})
}

func (u *UrlController) GetUrl(c *fiber.Ctx) error {
	id := c.Params("id")
	url, err := u.urlRepo.FindOne(id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON("Not Found")
	}
	return c.Redirect(url.OriginalUrl, http.StatusMovedPermanently)
}

func generateRandomCode(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		result[i] = charset[index.Int64()]
	}
	return string(result), nil
}

// apply testing all endpoints
// dockerize the app
// create github workflow for testing
// deploy on GCP / AWS ???
