package models

import "time"

type ProjectBlock struct {
	ID          uint                   `json:"id" gorm:"primaryKey;autoIncrement"`
	ProjectID   uint                   `json:"project_id" gorm:"not null"`
	Project     Project                `json:"project,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	BlockTypeID uint                   `json:"block_type_id" gorm:"not null"`
	BlockType   BlockType              `json:"block_type,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Position    int                    `json:"position" gorm:"default:0"`
	Data        map[string]interface{} `json:"data" gorm:"type:jsonb;default:'{}'"`
	CreatedAt   time.Time              `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time              `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   *time.Time             `json:"deleted_at,omitempty" gorm:"index"`
}
