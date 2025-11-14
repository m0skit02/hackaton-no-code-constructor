package repository

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	models "hackaton-no-code-constructor/pkg/model"
)

type ProjectPostgres struct {
	db *gorm.DB
}

func NewProjectRepo(db *gorm.DB) *ProjectPostgres { return &ProjectPostgres{db: db} }

func (r *ProjectPostgres) Create(project models.Project) (*models.Project, error) {
	logrus.Infof("[ProjectRepo] Creating project %+v", project)
	if err := r.db.Create(&project).Error; err != nil {
		logrus.Errorf("Error creating project %+v: %v", project, err)
		return nil, err
	}

	logrus.Infof("[ProjectRepo] Created project %+v", project)
	return &project, nil
}

func (r *ProjectPostgres) GetAll() ([]models.Project, error) {
	logrus.Infof("[ProjectRepo] Getting all projects")
	var projects []models.Project
	if err := r.db.Find(&projects).Error; err != nil {
		logrus.Errorf("[ProjectRepo] Error getting all projects: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectRepo] Getting all projects %+v", projects)
	return projects, nil
}

func (r *ProjectPostgres) GetByID(id uuid.UUID) (*models.Project, error) {
	logrus.Infof("[ProjectRepo] Getting project %+v", id)
	var project models.Project
	if err := r.db.First(&project, id).Error; err != nil {
		if err := gorm.ErrRecordNotFound; err != nil {
			logrus.Errorf("[ProjectRepo] Error getting project %+v: %v", id, err)
			return nil, nil
		}
		logrus.Errorf("[ProjectRepo] Error getting project %+v: %v", id, err)
		return nil, err
	}
	logrus.Infof("[ProjectRepo] Getting project %+v", project)
	return &project, nil
}

func (r *ProjectPostgres) Update(project models.Project) (*models.Project, error) {
	logrus.Infof("[ProjectRepo] Updating project %+v", project)
	if err := r.db.Save(&project).Error; err != nil {
		logrus.Errorf("Error updating project %+v: %v", project, err)
		return nil, err
	}
	logrus.Infof("[ProjectRepo] Updated project %+v", project)
	return &project, nil
}

func (r *ProjectPostgres) Delete(id uuid.UUID) error {
	logrus.Infof("[ProjectRepo] Deleting project %+v", id)
	if err := r.db.First(&models.Project{}, id).Error; err != nil {
		logrus.Errorf("Error getting project %+v: %v", id, err)
		return err
	}

	logrus.Infof("[ProjectRepo] Deleted project %+v", id)
	return nil
}
