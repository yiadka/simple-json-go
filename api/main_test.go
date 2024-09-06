package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.Default()
	r.GET("/albums", getAlbums)

	req, err := http.NewRequest(http.MethodGet, "/albums", nil)
	if err != nil {
		t.Fatalf("Failed to create a request: %v", err)
	}

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedBody := `[{"id":"1", "title":"Blue Train", "artist":"John Coltrane", "price": 56.99}, {"id":"2", "title":"Jeru", "artist":"Gerry Mulligan", "price":17.99}, {"id":"3", "title":"Sarah Vaughan and Clifford Brown", "artist": "","price":39.99}]`

	assert.JSONEq(t, expectedBody, w.Body.String())
}

func TestPostAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)

	request_body := `{"ID":"4", "Title":"hoge", "Artist":"fuga", "Price": 10.00}`

	r := gin.Default()
	r.POST("/albums", postAlbums)

	req, err := http.NewRequest(http.MethodPost, "/albums", bytes.NewBufferString(request_body))
	if err != nil {
		t.Fatalf("Failed to create a request: %v", err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	expectedBody := `{"id":"4", "title":"hoge", "artist":"fuga", "price": 10.00}`

	assert.JSONEq(t, expectedBody, w.Body.String())

	assert.Equal(t, 4, len(albums))
	assert.Equal(t, "4", albums[3].ID)
	assert.Equal(t, "hoge", albums[3].Title)
	assert.Equal(t, "fuga", albums[3].Artist)
	assert.Equal(t, 10.00, albums[3].Price)
}
