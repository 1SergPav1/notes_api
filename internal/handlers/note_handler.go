// Обработка CRUD-операций с заметками.
package handlers

import (
	"net/http"

	"github.com/1SergPav1/notes_api/internal/service"
	"github.com/gin-gonic/gin"
)

type NoteHandler struct {
	NoteService *service.NoteService
}

func NewNoteHandler(noteService *service.NoteService) *NoteHandler {
	return &NoteHandler{noteService}
}

// Создание заметки
func (h *NoteHandler) CreateNote(c *gin.Context) {
	var req struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		UserID uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
		return
	}

	err := h.NoteService.CreateNote(req.Title, req.Body, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Заметка создана"})
}
