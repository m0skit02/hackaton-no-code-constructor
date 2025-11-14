package repository

import (
	models "hackaton-no-code-constructor/pkg/model"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthPostgres { return &AuthPostgres{db: db} }

func (r *AuthPostgres) GetByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
