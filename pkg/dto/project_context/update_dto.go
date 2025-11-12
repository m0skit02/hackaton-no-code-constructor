package project_context

import "gorm.io/datatypes"

type UpdateProjectInput struct {
	Name string            `json:"name" binding:"omitempty,min=2"`
	Data datatypes.JSONMap `json:"data" binding:"omitempty"`
}
