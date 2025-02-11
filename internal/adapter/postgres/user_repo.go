package postgres

import (
	"github.com/1SergPav1/notes_api/internal/adapter"
	"github.com/1SergPav1/notes_api/internal/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) adapter.UserRepositiry {
	return &UserRepo{db}
}

// CreateUser implements adapter.UserRepositiry.
func (r *UserRepo) CreateUser(user *entity.User) error {
	return r.DB.Create(user).Error
}

// GetUserByUsername implements adapter.UserRepositiry.
func (r *UserRepo) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	return &user, err
}
