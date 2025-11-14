package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"errors"
	"hackaton-no-code-constructor/pkg/auth"
	"hackaton-no-code-constructor/pkg/dto/auth_context"
	"net/http"
)

func (h *Handler) Login(c *gin.Context) {
	var input auth_context.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.services.Auth.Login(input.Username, input.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := auth.GenerateToken(user.ID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
