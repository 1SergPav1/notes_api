package adapter

import "github.com/1SergPav1/notes_api/internal/entity"

type UserRepositiry interface {
	CreateUser(user *entity.User) error
	GetUserByUsername(username string) (*entity.User, error)
}

type NoteRepository interface {
	CreateNote(note *entity.Note) error
	GetNotesByUser(UserID uint) ([]entity.Note, error)
	UpdateNote(note *entity.Note) error
	DeleteNote(noteID uint) error
}
