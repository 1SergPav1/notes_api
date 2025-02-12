// Регистрация и авторизация пользователей.
package service

import (
	"errors"

	"github.com/1SergPav1/notes_api/internal/adapter"
	"github.com/1SergPav1/notes_api/internal/entity"
	"github.com/1SergPav1/notes_api/internal/utils"
)

type AuthService struct {
	UserRepo adapter.UserRepositiry
}

// Возвращает экземпляр сервиса авторизации
func NewAuthService(repo adapter.UserRepositiry) *AuthService {
	return &AuthService{repo}
}

// Логика регистрации пользователя
func (s *AuthService) Register(username, password string) error {
	existUser, _ := s.UserRepo.GetUserByUsername(username)
	if existUser != nil {
		return errors.New("Пользователь с таким именем уже существует")
	}

	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := entity.User{
		Username: username,
		Password: hashPassword,
	}

	return s.UserRepo.CreateUser(&user)
}

// Логика авторизации пользователя
func (s *AuthService) Login(username, password string) (string, error) {
	user, err := s.UserRepo.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("пользователь не найден")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return "", errors.New("неверный пароль")
	}

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}
