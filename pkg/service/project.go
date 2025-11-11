package service

import (
	"hackaton-no-code-constructor/pkg/repository"
)

type ProjectService struct {
	repo repository.Project
}

func NewProjectService(repo repository.Project) *ProjectService {
	return &ProjectService{repo: repo}
}

func (s *ProjectService) Create() {}

func (s *ProjectService) GetAll() {}

func (s *ProjectService) GetByID() {}

func (s *ProjectService) Update() {}

func (s *ProjectService) Delete() {}
