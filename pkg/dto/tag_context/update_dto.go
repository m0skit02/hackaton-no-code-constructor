package tag_context

type UpdateTagInput struct {
	Name string `json:"name" binding:"omitempty,min=2"`
}
