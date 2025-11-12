package models

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID        uuid.UUID              `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID    uuid.UUID              `json:"user_id" gorm:"not null"`
	Name      string                 `json:"name" gorm:"type:text;not null"`
	Data      map[string]interface{} `json:"data" gorm:"type:jsonb;default:'{}'"`
	CreatedAt time.Time              `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time              `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time             `json:"deleted_at,omitempty" gorm:"index"`
	Blocks    []ProjectBlock         `json:"blocks,omitempty" gorm:"foreignKey:ProjectID"`
}
