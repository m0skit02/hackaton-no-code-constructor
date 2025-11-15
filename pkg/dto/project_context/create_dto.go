package project_context

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type CreateProjectInput struct {
	UserID uuid.UUID      `json:"user_id" binding:"required"`
	Name   string         `json:"name" binding:"required,min=2"`
	Data   datatypes.JSON `json:"data" binding:"omitempty"`
}
