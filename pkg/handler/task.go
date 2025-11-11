package handler

import (
	models "aeza-checker/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	if len(input.Checks) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "checks cannot be empty"})
		return
	}

	created, err := h.services.Task.Create(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *Handler) getAllTask(c *gin.Context) {
	tasks, err := h.services.Task.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *Handler) updateTask(c *gin.Context) {
	id := c.Param("id")

	var input models.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	input.ID = models.MustParseUUID(id)

	updated, err := h.services.Task.Update(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

func (h *Handler) deleteTask(c *gin.Context) {
	id := c.Param("id")

	err := h.services.Task.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "task deleted"})
}
