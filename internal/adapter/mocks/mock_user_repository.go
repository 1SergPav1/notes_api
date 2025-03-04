package mocks

import (
	"errors"

	"github.com/1SergPav1/notes_api/internal/entity"
)

type MockUserRepository struct {
	Users map[string]entity.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		Users: make(map[string]entity.User),
	}
}

func (m *MockUserRepository) CreateUser(user *entity.User) error {
	if _, exists := m.Users[user.Username]; exists {
		return errors.New("Пользователь с таким именем уже существует")
	}

	m.Users[user.Username] = *user
	return nil
}

func (m *MockUserRepository) GetUserByUsername(username string) (*entity.User, error) {
	user, exists := m.Users[username]
	if !exists {
		return &entity.User{}, errors.New("такой пользователь не существует")
	}
	return &user, nil
}
