package project_block_context

import "github.com/google/uuid"

type CreateProjectBlockInput struct {
	ProjectID   uuid.UUID              `json:"project_id" binding:"required"`
	BlockTypeID uuid.UUID              `json:"block_type_id" binding:"required"`
	Position    int                    `json:"position" binding:"omitempty,gte=0"`
	Data        map[string]interface{} `json:"data" binding:"omitempty"`
}
