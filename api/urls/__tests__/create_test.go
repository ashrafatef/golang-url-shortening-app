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
)

func TestCreate(t *testing.T) {
	app := server.SetupServer()
	input := &urls.UrlInput{
		OriginalUrl: "https://www.gogle.com",
	}
	body, _ := json.Marshal(input)

	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/urls", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	res, err := app.Test(req, -1)

	if res.StatusCode != 200 {
		t.Error("Failed", err)
	}
	bodyA, _ := io.ReadAll(res.Body)
	
	var createdUrl repositories.Urls
	json.Unmarshal(bodyA, &createdUrl)

	if createdUrl.ShortUrl == "" {
		t.Error("Failed no short url")
	}

}

func TestHealth(t *testing.T) {
	app := server.SetupServer()
	req := httptest.NewRequest("GET", "/health", nil)
	fmt.Println(req.URL)
	res, _ := app.Test(req, -1)
	fmt.Println(res.StatusCode)

	if res.StatusCode != 200 {
		t.Error("Failed")
	}

}
