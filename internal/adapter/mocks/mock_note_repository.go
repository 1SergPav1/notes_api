package mocks

import (
	"errors"

	"github.com/1SergPav1/notes_api/internal/entity"
)

type MockNoteRepository struct {
	Notes map[uint]entity.Note
}

func NewMockNoteRepository() *MockNoteRepository {
	return &MockNoteRepository{
		Notes: make(map[uint]entity.Note),
	}
}

func (m *MockNoteRepository) CreateNote(note *entity.Note) error {
	note.ID = uint(len(m.Notes) + 1)
	m.Notes[note.ID] = *note
	return nil
}

func (m *MockNoteRepository) GetNotesByUser(UserID uint) ([]entity.Note, error) {
	var notes []entity.Note
	for _, note := range m.Notes {
		if note.UserID == UserID {
			notes = append(notes, note)
		}
	}
	return notes, nil
}

func (m *MockNoteRepository) UpdateNote(note *entity.Note) error {
	n, exists := m.Notes[note.ID]
	if !exists {
		return errors.New("Заметка не найдена")
	}
	n.Title = note.Title
	n.Body = note.Body
	m.Notes[note.ID] = n
	return nil
}

func (m *MockNoteRepository) DeleteNote(noteID uint) error {
	if _, exists := m.Notes[noteID]; !exists {
		return errors.New("Заметка не найдена")
	}
	delete(m.Notes, noteID)
	return nil
}
