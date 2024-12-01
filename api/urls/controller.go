package urls

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/ashrafatef/urlshortening/errors"
	"github.com/ashrafatef/urlshortening/repositories"
	"github.com/ashrafatef/urlshortening/validations"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/mattheath/base62"
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
	
	hashedUrl := hashUrl(input.OriginalUrl)

	id, error := u.urlRepo.Create(repositories.UrlInput{
		Url:       input.OriginalUrl,
		HashedUrl: hashedUrl,
	})

	if error != nil {
		logrus.Error("Error creating url: ", error)
		return c.Status(http.StatusInternalServerError).JSON("Internal Server Error")
	}
	return c.JSON(id)
}

func (u *UrlController) GetUrl(c *fiber.Ctx) error {
	fmt.Println("NOT HEEEEEERE")
	id := c.Params("id")
	fmt.Printf("%+v\n", id)
	url, err := u.urlRepo.FindOne(id)
	fmt.Printf("yrl is %+v\n", url)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON("Not Found")
	}
	return c.Redirect(url.OriginalUrl, http.StatusMovedPermanently)
}

func (u *UrlController) GetUrls(c *fiber.Ctx) error {
	fmt.Println(" HEEEEEERE")
	return c.JSON("Urls")
}

func hashUrl(url string) string {
	bigInt := new(big.Int)

	hashed := sha256.Sum256([]byte(url + string(time.Now().UnixMilli())))
	hashedUrl := bigInt.SetBytes(hashed[:])

	encoded := base62.EncodeBigInt(hashedUrl)
	fmt.Printf("encoded is %+v\n", string(encoded))
	return string(encoded)[0:8]
}
