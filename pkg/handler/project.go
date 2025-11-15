package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hackaton-no-code-constructor/pkg/dto/project_context"
	"net/http"
)

func (h *Handler) createProject(c *gin.Context) {
	var input project_context.CreateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Errorf("[ProjectHandler] Binding error %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProject, err := h.services.Project.CreateProject(input)
	if err != nil {
		logrus.Errorf("[ProjectHandler] CreateProject error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectHandler] CreateProject success: %v", createdProject)
	c.JSON(http.StatusOK, gin.H{"project": createdProject})
}

func (h *Handler) getAllProject(c *gin.Context) {
	projects, err := h.services.Project.GetAllProjects()
	if err != nil {
		logrus.Errorf("[ProjectHandler] GetAllProjects error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectHandler] GetAllProjects success: %v", projects)
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func (h *Handler) getByIdProject(c *gin.Context) {
	id := c.Param("id")

	project, err := h.services.Project.GetByIDProject(parseUUID(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("[ProjectHandler] GetByIdProject error %+v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		logrus.Errorf("[ProjectHandler] GetByIdProject error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectHandler] GetByIdProject success: %v", project)
	c.JSON(http.StatusOK, gin.H{"project": project})
}

func (h *Handler) getMyProjects(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	projects, err := h.services.Project.GetByUserID(userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, projects)
}

func (h *Handler) updateProject(c *gin.Context) {
	id := c.Param("id")

	var input project_context.UpdateProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		logrus.Errorf("[ProjectHandler] Binding error %+v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProject, err := h.services.Project.UpdateProject(parseUUID(id), input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Errorf("[ProjectHandler] UpdateProject error %+v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		logrus.Errorf("[ProjectHandler] UpdateProject error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectHandler] UpdateProject success: %v", updatedProject)
	c.JSON(http.StatusOK, gin.H{"project": updatedProject})
}

func (h *Handler) deleteProject(c *gin.Context) {
	id := c.Param("id")

	if err := h.services.Project.DeleteProject(parseUUID(id)); err != nil {
		logrus.Errorf("[ProjectHandler] DeleteProject error %+v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.Infof("[ProjectHandler] DeleteProject success: %v", id)
	c.Status(http.StatusNoContent)
}
