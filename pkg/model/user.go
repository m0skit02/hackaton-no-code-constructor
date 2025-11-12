package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name         string     `json:"name" gorm:"type:varchar(255);not null"`
	Username     string     `json:"username" gorm:"type:varchar(255);unique;not null"`
	PasswordHash string     `json:"password_hash" gorm:"type:varchar(255);not null"`
	CreatedAt    time.Time  `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty" gorm:"index"`
	Projects     []Project  `json:"projects,omitempty" gorm:"foreignKey:UserID"`
}
