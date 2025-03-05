package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1SergPav1/notes_api/internal/adapter/mocks"
	"github.com/1SergPav1/notes_api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockRepoUser = mocks.NewMockUserRepository()

func TestRegister(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authService := service.NewAuthService(mockRepoUser)
	authHandler := NewAuthHandler(authService)

	router := gin.Default()
	router.POST("/auth/register", authHandler.Register)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"password": "test",
	})

	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Пользователь зарегистрирован")
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authService := service.NewAuthService(mockRepoUser)
	authHandler := NewAuthHandler(authService)

	router := gin.Default()
	router.POST("/auth/login", authHandler.Login)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"password": "test",
	})

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "token")
}

func TestLiginInvalidPassword(t *testing.T) {
	gin.SetMode(gin.TestMode)

	authService := service.NewAuthService(mockRepoUser)
	authHandler := NewAuthHandler(authService)

	router := gin.Default()
	router.POST("/auth/login", authHandler.Login)

	reqBody, _ := json.Marshal(map[string]string{
		"username": "testuser",
		"password": "wrongpassword",
	})

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "неверный пароль")
}
