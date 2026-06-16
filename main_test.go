package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthHandler(t *testing.T) {

	gin.SetMode(gin.TestMode)

	router := gin.Default()
	router.GET("/health", HealthHandler)

	req, _ := http.NewRequest("GET", "/health", nil)

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status 200 but got %d", w.Code)
	}
}
