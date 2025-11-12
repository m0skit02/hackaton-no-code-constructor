package models

import (
	"time"

	"github.com/google/uuid"
)

type ProjectBlock struct {
	ID          uuid.UUID              `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ProjectID   uuid.UUID              `json:"project_id" gorm:"not null"`
	BlockTypeID uuid.UUID              `json:"block_type_id" gorm:"not null"`
	Position    int                    `json:"position" gorm:"default:0"`
	Data        map[string]interface{} `json:"data" gorm:"type:jsonb;default:'{}'"`
	CreatedAt   time.Time              `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time              `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time             `json:"deleted_at,omitempty" gorm:"index"`
}
