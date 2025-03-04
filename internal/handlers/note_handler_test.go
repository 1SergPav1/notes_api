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

func TestCreateNote(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockRepo := mocks.NewMockNoteRepository()
	noteService := service.NewNoteService(mockRepo)
	noteHandler := NewNoteHandler(noteService)

	router := gin.Default()
	router.POST("/notes/", noteHandler.CreateNote)

	reqBody, _ := json.Marshal(map[string]interface{}{
		"title":   "test note",
		"body":    "testing test of tests",
		"user_id": 1,
	})

	req, _ := http.NewRequest("POST", "/notes/", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-type", "application/json")
	req.Header.Set("Authorization", "Bearer valid_token")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)
	assert.Contains(t, recorder.Body.String(), "Заметка создана")
}
