package block_type_context

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type UpdateBlockTypeInput struct {
	TagID       uuid.UUID         `json:"tag_id" binding:"required"`
	Name        string            `json:"name" binding:"required,min=2"`
	Description string            `json:"description" binding:"omitempty"`
	Template    string            `json:"template" binding:"required"`
	Schema      datatypes.JSONMap `json:"schema" binding:"required"`
	Preview     string            `json:"preview" binding:"omitempty,url"`
}
