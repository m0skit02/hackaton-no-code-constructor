package models

import "time"

type Tag struct {
	ID         uint        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string      `json:"name" gorm:"type:text;not null"`
	CreatedAt  time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  time.Time   `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  *time.Time  `json:"deleted_at,omitempty" gorm:"index"`
	BlockTypes []BlockType `json:"block_types,omitempty" gorm:"foreignKey:TagID"`
}
