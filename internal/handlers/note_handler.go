// Обработка CRUD-операций с заметками.
package handlers

import (
	"net/http"
	"strconv"

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

// Получение списка заметок пользователя
func (h *NoteHandler) GetNotes(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "Некорректный ID пользователя"})
		return
	}

	notes, err := h.NoteService.NoteRepo.GetNotesByUser(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, notes)
}

// Обновление заметки
func (h *NoteHandler) UpdateNote(c *gin.Context) {
	noteID, err := strconv.ParseUint(c.Query("note_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID заметки"})
		return
	}

	var req struct {
		Title  string `json:"title"`
		Body   string `json:"body"`
		UserID uint   `json:"user_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректные данные"})
	}

	err = h.NoteService.UpdateNote(uint(noteID), req.Title, req.Body, req.UserID)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заметка обновлена"})
}

// Удаление заметки
func (h *NoteHandler) DeleteNote(c *gin.Context) {
	noteID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID заметки"})
		return
	}

	userID, err := strconv.ParseUint(c.Query("user_id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Некорректный ID пользователя"})
		return
	}

	err = h.NoteService.DeleteNote(uint(noteID), uint(userID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Заметка удалена"})
}
