// Генерация и валидация токенов.
package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ReadToken(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка открытия файла с токеном: %w", err)
	}

	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("ошибка чтения файла с токеном: %w", err)
	}

	token := strings.TrimSpace(string(data))
	return token, nil
}

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 12).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	key, err := ReadToken("jwt_token.txt")
	if err != nil {
		return "", errors.New("не удалось считать ключ для создания jwt-токена")
	}
	return token.SignedString([]byte(key))
}
