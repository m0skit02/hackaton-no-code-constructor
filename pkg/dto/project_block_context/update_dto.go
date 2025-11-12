package project_block_context

import (
	"github.com/google/uuid"
)

type UpdateProjectBlockInput struct {
	BlockTypeID *uuid.UUID              `json:"block_type_id,omitempty" binding:"omitempty"`
	Position    *int                    `json:"position,omitempty" binding:"omitempty,gte=0"`
	Data        *map[string]interface{} `json:"data,omitempty" binding:"omitempty"`
}
