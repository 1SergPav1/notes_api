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

func (r *NoteRepo) CreateNote(note *entity.Note) error {
	return r.DB.Create(note).Error
}

func (r *NoteRepo) DeleteNote(noteID uint) error {
	return r.DB.Delete(&entity.Note{}, noteID).Error
}

func (r *NoteRepo) GetNotesByUser(UserID uint) ([]entity.Note, error) {
	var notes []entity.Note
	result := r.DB.Where("user_id = ?", UserID).Find(&notes)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return notes, nil
}

func (r *NoteRepo) UpdateNote(note *entity.Note) error {
	return r.DB.Save(note).Error
}
