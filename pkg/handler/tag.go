package handler

import (
	"errors"
	"hackaton-no-code-constructor/pkg/dto/tag_context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) createTag(c *gin.Context) {
	var input tag_context.CreateTagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTag, err := h.services.Tag.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, createdTag)
}

func (h *Handler) getAllTags(c *gin.Context) {
	tags, err := h.services.Tag.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, tags)
}

func (h *Handler) getTagById(c *gin.Context) {
	id := c.Param("id")

	tag, err := h.services.Tag.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func (h *Handler) updateTag(c *gin.Context) {
	id := c.Param("id")

	var input tag_context.UpdateTagInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	updatedTag, err := h.services.Tag.Update(id, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, updatedTag)
}

func (h *Handler) deleteTag(c *gin.Context) {
	id := c.Param("id")

	if err := h.services.Tag.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{})
}
