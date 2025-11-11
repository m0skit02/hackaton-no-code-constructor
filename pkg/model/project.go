package models

import "time"

type Project struct {
	ID        uint                   `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint                   `json:"user_id" gorm:"not null"`
	User      User                   `json:"user,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Name      string                 `json:"name" gorm:"type:text;not null"`
	Data      map[string]interface{} `json:"data" gorm:"type:jsonb;default:'{}'"`
	CreatedAt time.Time              `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time              `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time             `json:"deleted_at,omitempty" gorm:"index"`
	Blocks    []ProjectBlock         `json:"blocks,omitempty" gorm:"foreignKey:ProjectID"`
}
