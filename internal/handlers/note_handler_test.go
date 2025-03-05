package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/1SergPav1/notes_api/internal/adapter/mocks"
	"github.com/1SergPav1/notes_api/internal/middleware"
	"github.com/1SergPav1/notes_api/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var mockRepoNote = mocks.NewMockNoteRepository()

func TestCreateNote(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteService := service.NewNoteService(mockRepoNote)
	noteHandler := NewNoteHandler(noteService)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.POST("/notes/", noteHandler.CreateNote)

	reqBody, _ := json.Marshal(map[string]any{
		"title":   "test note",
		"body":    "testing test of tests",
		"user_id": 1,
	})

	req, _ := http.NewRequest("POST", "/notes/", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJleHAiOjE3NDEyMDc5OTUsImlhdCI6MTc0MTE2NDc5NX0.IlJ3u-Y-OyXcnA3f1U35dRTivmEbrBhIuCt0llFURM8")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Заметка создана")
}

func TestUpdateNote(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteService := service.NewNoteService(mockRepoNote)
	noteHandler := NewNoteHandler(noteService)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.PUT("/notes/:id", noteHandler.UpdateNote)

	reqBody, _ := json.Marshal(map[string]any{
		"user_id": 1,
		"title":   "new test note",
		"body":    "new testing test of tests",
	})

	req, _ := http.NewRequest("PUT", "/notes/1?user_id=1", bytes.NewBuffer(reqBody))
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJleHAiOjE3NDEyMDc5OTUsImlhdCI6MTc0MTE2NDc5NX0.IlJ3u-Y-OyXcnA3f1U35dRTivmEbrBhIuCt0llFURM8")
	req.Header.Set("Content-type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Заметка обновлена")
}

func TestGetNotes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteService := service.NewNoteService(mockRepoNote)
	noteHandler := NewNoteHandler(noteService)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.GET("/notes/", noteHandler.GetNotes)

	req, _ := http.NewRequest("GET", "/notes/?user_id=1", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJleHAiOjE3NDEyMDc5OTUsImlhdCI6MTc0MTE2NDc5NX0.IlJ3u-Y-OyXcnA3f1U35dRTivmEbrBhIuCt0llFURM8")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "new test note")
}

func TestDeleteNoteForeign(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteService := service.NewNoteService(mockRepoNote)
	noteHandler := NewNoteHandler(noteService)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.DELETE("/notes/:id", noteHandler.DeleteNote)

	req, _ := http.NewRequest("DELETE", "/notes/1?user_id=2", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJleHAiOjE3NDEyMDc5OTUsImlhdCI6MTc0MTE2NDc5NX0.IlJ3u-Y-OyXcnA3f1U35dRTivmEbrBhIuCt0llFURM8")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusForbidden, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Заметка не найдена")
}

func DeleteNote(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteService := service.NewNoteService(mockRepoNote)
	noteHandler := NewNoteHandler(noteService)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.DELETE("/notes/:id", noteHandler.DeleteNote)

	req, _ := http.NewRequest("DELETE", "/note/1?user_id=1", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJleHAiOjE3NDEyMDc5OTUsImlhdCI6MTc0MTE2NDc5NX0.IlJ3u-Y-OyXcnA3f1U35dRTivmEbrBhIuCt0llFURM8")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Заметка удалена")
}

func TestCreateNoteWithoutToken(t *testing.T) {
	gin.SetMode(gin.TestMode)

	noteService := service.NewNoteService(mockRepoNote)
	noteHandler := NewNoteHandler(noteService)

	router := gin.Default()
	router.Use(middleware.AuthMiddleware())
	router.POST("/notes/", noteHandler.CreateNote)

	reqBody, _ := json.Marshal(map[string]any{
		"title":   "test note",
		"body":    "testing test of tests",
		"user_id": 1,
	})

	req, _ := http.NewRequest("POST", "/notes/", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "токен не найден")
}
