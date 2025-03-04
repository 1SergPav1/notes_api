// Генерация и валидация токенов.
package utils

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func ReadToken(filename string) (string, error) {
	_, dir, _, _ := runtime.Caller(1)
	keyPath := filepath.Join(filepath.Dir(dir), "../..", filename)

	file, err := os.Open(keyPath)
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

func GenerateJWT(user_id int64) (string, error) {
	claims := Claims{
		UserID: user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, err := ReadToken("./jwt_token.txt")
	if err != nil {
		return "", errors.New("не удалось считать ключ для создания jwt-токена")
	}
	return token.SignedString([]byte(key))
}

func ParseJWT(tokenString string) (*Claims, error) {
	secretKey, err := ReadToken("./jwt_token.txt")
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("токен не действителен")
	}

	return claims, nil
}
