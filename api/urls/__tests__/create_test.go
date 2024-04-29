package urls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashrafatef/urlshortening/api/urls"
	"github.com/ashrafatef/urlshortening/server"
	"github.com/gofiber/fiber/v2/log"
)

func TestCreate(t *testing.T) {
	app := server.SetupServer()
	input := &urls.UrlInput{
		OriginalUrl: "https://www.gogle.com",
	}
	body, _ := json.Marshal(input)
	fmt.Println(bytes.NewReader(body))
	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000/urls", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")


	res, err := app.Test(req, -1)
	
	if res.StatusCode != 200 {
		t.Error("Failed", err)
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

// func TestCreate(t *testing.T) {
// 	app := server.SetupServer()
// 	server := httptest.NewServer(app.)
// 	defer server.Close()
// 	input := &urls.UrlInput{
// 		OriginalUrl: "https://www.gogle.com",
// 	}
// 	body, _ := json.Marshal(input)

// 	req := httptest.NewRequest(http.MethodPost, "http://localhost:3000", bytes.NewReader(body))
// 	defer req.Body.Close()
// 	fmt.Println(req.URL)
// 	res, _ := app.Test(req, -1)

// 	if res.StatusCode != 200 {
// 		t.Error("Failed")
// 	}

// }
