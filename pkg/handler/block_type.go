package handler

import (
	"errors"
	"hackaton-no-code-constructor/pkg/dto/block_type_context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *Handler) createBlockType(c *gin.Context) {
	var input block_type_context.CreateBlockTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBlockType, err := h.services.BlockType.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdBlockType)
}

func (h *Handler) getAllBlockTypes(c *gin.Context) {
	blockTypes, err := h.services.BlockType.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, blockTypes)
}

func (h *Handler) getBlockTypeById(c *gin.Context) {
	id := c.Param("id")

	blockType, err := h.services.BlockType.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Block type not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blockType)
}

func (h *Handler) updateBlockType(c *gin.Context) {
	id := c.Param("id")

	var input block_type_context.UpdateBlockTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedBlockType, err := h.services.BlockType.Update(id, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Block type not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBlockType)
}

func (h *Handler) deleteBlockType(c *gin.Context) {
	id := c.Param("id")

	if err := h.services.BlockType.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
}
