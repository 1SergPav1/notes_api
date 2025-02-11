package postgres

import (
	"github.com/1SergPav1/notes_api/internal/adapter"
	"github.com/1SergPav1/notes_api/internal/entity"
	"gorm.io/gorm"
)

type NoteRepo struct {
	DB *gorm.DB
}

func NewNoteRepo(db *gorm.DB) adapter.NoteRepository {
	return &NoteRepo{db}
}

// CreateNote implements adapter.NoteRepository.
func (r *NoteRepo) CreateNote(note *entity.Note) error {
	return r.DB.Create(note).Error
}

// DeleteNote implements adapter.NoteRepository.
func (r *NoteRepo) DeleteNote(noteID uint) error {
	return r.DB.Delete(&entity.Note{}, noteID).Error
}

// GetNotesByUser implements adapter.NoteRepository.
func (r *NoteRepo) GetNotesByUser(UserID uint) ([]entity.Note, error) {
	var notes []entity.Note
	result := r.DB.Where("user_id = ?", UserID).Find(&notes)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}

// UpdateNote implements adapter.NoteRepository.
func (r *NoteRepo) UpdateNote(note *entity.Note) error {
	return r.DB.Save(note).Error
}
