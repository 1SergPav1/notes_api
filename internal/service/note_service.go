// Работа с заметками: создание, редактирование, удаление.
package service

import (
	"errors"

	"github.com/1SergPav1/notes_api/internal/adapter"
	"github.com/1SergPav1/notes_api/internal/entity"
)

type NoteService struct {
	NoteRepo adapter.NoteRepository
}

func NewNoteService(repo adapter.NoteRepository) *NoteService {
	return &NoteService{repo}
}

// Создание заметки
func (s *NoteService) CreateNote(title, body string, userID uint) error {
	if body == "" || title == "" {
		return errors.New("название и текст заметки не могуть быть пустыми")
	}

	note := entity.Note{
		Title:  title,
		Body:   body,
		UserID: userID,
	}

	return s.NoteRepo.CreateNote(&note)
}

// Получение всех заметок пользователя
func (s *NoteService) GetNotes(userID uint) ([]entity.Note, error) {
	return s.NoteRepo.GetNotesByUser(userID)
}

// Обновление заметки
func (s *NoteService) UpdateNote(noteID uint, title, body string, userID uint) error {
	notes, err := s.NoteRepo.GetNotesByUser(userID)
	if err != nil {
		return err
	}

	var noteToUpdate *entity.Note
	for _, note := range notes {
		if note.ID == noteID {
			noteToUpdate = &note
			break
		}
	}

	if noteToUpdate == nil {
		return errors.New("Заметка не найдена")
	}

	noteToUpdate.Title = title
	noteToUpdate.Body = body

	return s.NoteRepo.UpdateNote(noteToUpdate)
}

// Удаление заметки
func (s *NoteService) DeleteNote(noteID, userID uint) error {
	notes, err := s.NoteRepo.GetNotesByUser(userID)
	if err != nil {
		return err
	}

	var noteToDel *entity.Note
	for _, note := range notes {
		if note.ID == noteID {
			noteToDel = &note
			break
		}
	}

	if noteToDel == nil {
		return errors.New("Заметка не найдена")
	}

	return s.NoteRepo.DeleteNote(noteToDel.ID)
}
