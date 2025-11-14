package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hackaton-no-code-constructor/pkg/dto/project_block_context"
	"net/http"
)

func (h *Handler) createProjectBlock(c *gin.Context) {
	var input project_block_context.CreateProjectBlockInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Errorf("[ProjectBlockHandler] Binding input err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProjectBlock, err := h.services.ProjectBlock.CreateProjectBlock(input)
	if err != nil {
		logrus.Errorf("[ProjectBlockHandler] CreateProjectBlock err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectBlockHandler] CreateProjectBlock success: %v", createdProjectBlock)
	c.JSON(http.StatusOK, gin.H{"projectBlock": createdProjectBlock})
}

func (h *Handler) getAllProjectBlock(c *gin.Context) {
	users, err := h.services.ProjectBlock.GetAllProjectBlock()
	if err != nil {
		logrus.Errorf("[ProjectBlockHandler] GetAllProjectBlock err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectBlockHandler] GetAllProjectBlock success: %v", users)
	c.JSON(http.StatusOK, gin.H{"projectBlocks": users})
}

func (h *Handler) getByIdProjectBlock(c *gin.Context) {
	id := c.Param("id")

	projectBlock, err := h.services.ProjectBlock.GetByIDProjectBlock(parseUUID(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("[ProjectBlockHandler] GetByIdProjectBlock err: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logrus.Errorf("[ProjectBlockHandler] GetByIdProjectBlock err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectBlockHandler] GetByIdProjectBlock success: %v", projectBlock)
	c.JSON(http.StatusOK, gin.H{"projectBlock": projectBlock})
}

func (h *Handler) updateProjectBlock(c *gin.Context) {
	id := c.Param("id")

	var input project_block_context.UpdateProjectBlockInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Errorf("[ProjectBlockHandler] Binding input err: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProjectBlock, err := h.services.ProjectBlock.UpdateProjectBlock(parseUUID(id), input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("[ProjectBlockHandler] UpdateProjectBlock err: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		logrus.Errorf("[ProjectBlockHandler] UpdateProjectBlock err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectBlockHandler] UpdateProjectBlock success: %v", updatedProjectBlock)
	c.JSON(http.StatusOK, gin.H{"projectBlock": updatedProjectBlock})
}

func (h *Handler) deleteProjectBlock(c *gin.Context) {
	id := c.Param("id")

	if err := h.services.ProjectBlock.DeleteProjectBlock(parseUUID(id)); err != nil {
		logrus.Errorf("[ProjectBlockHandler] DeleteProjectBlock err: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectBlockHandler] DeleteProjectBlock success: %v", id)
	c.JSON(http.StatusOK, gin.H{"projectBlock": nil})
}
