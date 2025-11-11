package models

import (
	"github.com/lib/pq"
	"time"
)

type BlockType struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	TagID         *uint          `json:"tag_id,omitempty"`
	Tag           *Tag           `json:"tag,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Name          string         `json:"name" gorm:"type:text;not null"`
	Description   string         `json:"description,omitempty" gorm:"type:text"`
	Template      string         `json:"template" gorm:"type:text;not null"`
	Schema        pq.StringArray `json:"schema" gorm:"type:jsonb;not null"`
	Preview       string         `json:"preview,omitempty" gorm:"type:text"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     *time.Time     `json:"deleted_at,omitempty" gorm:"index"`
	ProjectBlocks []ProjectBlock `json:"project_blocks,omitempty" gorm:"foreignKey:BlockTypeID"`
}
