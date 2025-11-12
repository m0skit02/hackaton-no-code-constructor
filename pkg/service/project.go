package service

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"hackaton-no-code-constructor/pkg/dto/project_context"
	models "hackaton-no-code-constructor/pkg/model"
	"hackaton-no-code-constructor/pkg/repository"
)

type ProjectService struct {
	repo repository.Project
}

func NewProjectService(repo repository.Project) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) CreateProject(input project_context.CreateProjectInput) (*models.Project, error) {
	logrus.Infof("[ProjectService] Create Project %+v", input)

	project := models.Project{
		Name: input.Name,
		Data: input.Data,
	}

	createdProject, err := s.repo.Create(project)
	if err != nil {
		logrus.Errorf("[ProjectService] Create Project err: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectService] Create Project success: %v", createdProject)
	return createdProject, nil
}

func (s *ProjectService) GetAllProjects() ([]models.Project, error) {
	logrus.Infof("[ProjectService] Get All Projects")
	projects, err := s.repo.GetAll()
	if err != nil {
		logrus.Errorf("[ProjectService] Get All Projects err: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectService] Get All Projects success: %v", projects)
	return projects, nil
}

func (s *ProjectService) GetByIDProject(id uuid.UUID) (*models.Project, error) {
	logrus.Infof("[ProjectService] Get By ID %v", id)

	project, err := s.repo.GetByID(id)
	if err != nil {
		logrus.Errorf("[ProjectService] Get By ID err: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectService] Get By ID success: %v", project)
	return project, nil
}

func (s *ProjectService) UpdateProject(id uuid.UUID, input project_context.UpdateProjectInput) (*models.Project, error) {
	logrus.Infof("[ProjectService] Update Project %+v", input)

	project, err := s.repo.GetByID(id)
	if err != nil {
		logrus.Errorf("[ProjectService] Update Project err: %v", err)
		return nil, err
	}

	if input.Name != "" {
		project.Name = input.Name
	}
	if input.Data != nil {
		project.Data = input.Data
	}

	updatedProject, err := s.repo.Update(*project)
	if err != nil {
		logrus.Errorf("[ProjectService] Update Project err: %v", err)
		return nil, err
	}

	logrus.Infof("[ProjectService] Update Project success: %v", updatedProject)
	return updatedProject, nil
}

func (s *ProjectService) DeleteProject(id uuid.UUID) error {
	logrus.Infof("[ProjectService] Delete Project %+v", id)
	if err := s.repo.Delete(id); err != nil {
		logrus.Errorf("[ProjectService] Delete Project err: %v", err)
		return err
	}

	logrus.Infof("[ProjectService] Delete Project success: %v", id)
	return nil
}
