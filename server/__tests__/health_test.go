package server

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/ashrafatef/urlshortening/server"
)

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
