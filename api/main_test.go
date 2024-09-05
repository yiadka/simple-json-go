package main

import (
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
