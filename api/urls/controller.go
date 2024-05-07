package urls

import (
	"crypto/md5"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ashrafatef/urlshortening/repositories"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/mattheath/base62"
)

type UrlInput struct {
	OriginalUrl string `json:"original_url"`
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
	fmt.Printf("%+v", c.Body())
	if err := c.BodyParser(&input); err != nil {
		log.Error(err)
		return err
	}
	fmt.Printf("%+v", input)
	hashedUrl := hashUrl(input.OriginalUrl)

	id := u.urlRepo.Create(repositories.UrlInput{
		Url:       input.OriginalUrl,
		HashedUrl: hashedUrl,
	})

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
	hasher := md5.New()
	bigInt := new(big.Int)
	hashedUrl := bigInt.SetBytes(hasher.Sum([]byte(url)))

	encoded := base62.EncodeBigInt(hashedUrl)
	fmt.Printf("encoded is %+v\n", string(encoded))
	return string(encoded)[0:8]
}
