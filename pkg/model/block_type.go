package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type BlockType struct {
	ID            uuid.UUID         `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	TagID         uuid.UUID         `json:"tag_id"`
	Name          string            `json:"name" gorm:"type:text;not null"`
	Description   string            `json:"description,omitempty" gorm:"type:text"`
	Template      string            `json:"template" gorm:"type:text;not null"`
	Schema        datatypes.JSONMap `json:"schema" gorm:"type:jsonb;not null"`
	Preview       string            `json:"preview,omitempty" gorm:"type:text"`
	CreatedAt     time.Time         `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time         `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     *time.Time        `json:"deleted_at,omitempty" gorm:"index"`
	ProjectBlocks []ProjectBlock    `json:"project_blocks,omitempty" gorm:"foreignKey:BlockTypeID"`
}
