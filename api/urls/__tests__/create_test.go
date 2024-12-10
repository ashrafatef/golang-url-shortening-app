package urls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashrafatef/urlshortening/api/urls"
	"github.com/ashrafatef/urlshortening/repositories"
	"github.com/ashrafatef/urlshortening/server"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	assert := assert.New(t)

	t.Run("Create url", func(t *testing.T) {
		app := server.SetupServer()
		input := &urls.UrlInput{
			OriginalUrl: "https://www.gogle.com",
		}
		body, _ := json.Marshal(input)

		req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/urls", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		res, _ := app.Test(req, -1)

		assert.Equal(200, res.StatusCode)
		bodyA, _ := io.ReadAll(res.Body)

		var createdUrl repositories.Urls
		json.Unmarshal(bodyA, &createdUrl)
		fmt.Println(bodyA)
		assert.Len(createdUrl.ShortUrl, 8)
	})

	t.Run("Create unique short for same two urls", func(t *testing.T) {
		app := server.SetupServer()
		input := []urls.UrlInput{
			{OriginalUrl: "https://www.google.com"},
			{OriginalUrl: "https://www.google.com"},
		}

		createShortUrl := func(url urls.UrlInput) string {
			body, _ := json.Marshal(url)

			req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/urls", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			res, _ := app.Test(req, -1)

			// assert.Equal(200, res.StatusCode)
			bodyA, _ := io.ReadAll(res.Body)

			var createdUrl repositories.Urls
			json.Unmarshal(bodyA, &createdUrl)

			return createdUrl.ShortUrl
		}

		firstUrl := createShortUrl(input[0])
		secondUrl := createShortUrl(input[1])
		fmt.Println(firstUrl)
		fmt.Println(secondUrl)
		assert.NotEqual(firstUrl, secondUrl)
	})
}
