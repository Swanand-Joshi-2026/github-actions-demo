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

	// Intentionally broken assertion
	if w.Code != http.StatusInternalServerError {
		t.Errorf("expected status 500 but got %d", w.Code)
	}
}
