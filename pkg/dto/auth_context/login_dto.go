package auth_context

type LoginInput struct {
	Username     string `json:"username" binding:"required"`
	PasswordHash string `json:"password_hash" binding:"required"`
}
