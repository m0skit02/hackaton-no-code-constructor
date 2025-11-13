package repository

import (
	models "hackaton-no-code-constructor/pkg/model"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthPostgres { return &AuthPostgres{db: db} }

func (r *AuthPostgres) GetUserByUsernameAndPasswordHash(username string, passwordHash string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "username = ? AND password_hash = ?", username, passwordHash).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
