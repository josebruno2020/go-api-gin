package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/josebruno2020/go-api-gin/routes"
	"github.com/stretchr/testify/assert"
)

func SetupGinInstance() *gin.Engine {
	r := gin.Default()
	return r
}

func TestHealthCheckRoute(t *testing.T) {
	r := SetupGinInstance()
	r.GET("/", routes.HealthCheck)

	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	r.ServeHTTP(res, req)

	assert.Equal(t, http.StatusOK, res.Code)

	mockResponse := `{"ping":"pong"}`

	resBody, err := io.ReadAll(res.Body)

	if err != nil {
		t.Fatalf("Erro: %s", err.Error())
	}

	assert.Equal(t, mockResponse, string(resBody))
}
