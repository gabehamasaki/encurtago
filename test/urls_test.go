package tests

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gabehamasaki/encurtago/internal/config"
	"github.com/gabehamasaki/encurtago/internal/database"
	"github.com/gabehamasaki/encurtago/internal/database/connection"
	"github.com/gabehamasaki/encurtago/internal/dtos"
	"github.com/gabehamasaki/encurtago/internal/router"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func setupTest() (*gin.Engine, *pgx.Conn) {
	r := gin.Default()

	cfg := config.NewConfig()

	conn, err := connection.NewConnection(context.Background(), cfg)
	if err != nil {
		panic(err)
		return nil, nil
	}

	db := database.New(conn)
	cfg.SetDB(db)

	router.RegisterRoutes(r, cfg)

	return r, conn
}

func TestCreateShortUrlRoute(t *testing.T) {
	r, conn := setupTest()
	defer conn.Close(context.Background())

	w := httptest.NewRecorder()

	exampleShortURl := &dtos.CreateShortURLRequest{
		Original: "https://www.google.com",
	}

	requestBody, _ := json.Marshal(exampleShortURl)

	req, _ := http.NewRequest("POST", "/api/urls", strings.NewReader(string(requestBody)))
	r.ServeHTTP(w, req)

	resBody := &dtos.CreateShortURLResponse{}
	json.Unmarshal(w.Body.Bytes(), resBody)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "https://www.google.com", resBody.Original)
}

type UrlsResponse struct {
	Urls []dtos.URL `json:"urls"`
}

func TestListURls(t *testing.T) {
	r, conn := setupTest()
	defer conn.Close(context.Background())

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/api/urls", nil)
	r.ServeHTTP(w, req)

	body := &UrlsResponse{}
	json.Unmarshal(w.Body.Bytes(), body)

	assert.Equal(t, 200, w.Code)
	assert.NotEmpty(t, body.Urls)
}
