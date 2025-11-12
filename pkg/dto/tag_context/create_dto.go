package tag_context

type CreateTagInput struct {
	Name string `json:"name" binding:"required,min=2"`
}
