package urls

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ashrafatef/urlshortening/db"
	"github.com/ashrafatef/urlshortening/repositories"
	"github.com/ashrafatef/urlshortening/server"
)

func TestGetUrl(t *testing.T) {
	app := server.SetupServer()
	dbConn := db.Connect()
	urlRepo := repositories.NewUrlRepository(dbConn)

	t.Run("redirects if the url found", func(t *testing.T) {
		createdUrl := urlRepo.Create(repositories.UrlInput{
			Url:       "https://www.gogle.com",
			HashedUrl: "x1x1x2",
		})
		urlWithQuery := "http://localhost:3000/urls/" + createdUrl.ShortUrl
		req := httptest.NewRequest(http.MethodGet, urlWithQuery, nil)
		req.Header.Set("Content-Type", "application/json")
	
		res, err := app.Test(req, -1)
	
		if res.StatusCode != 301 {
			t.Error("Failed", err)
		}
	})

	t.Run("returns 404 if not found", func(t *testing.T) {

		urlWithQuery := "http://localhost:3000/urls/xxx-not-found" 
		req := httptest.NewRequest(http.MethodGet, urlWithQuery, nil)
		req.Header.Set("Content-Type", "application/json")
	
		res, err := app.Test(req, -1)
	
		if res.StatusCode != 404 {
			t.Error("Failed", err)
		}
	})

}
