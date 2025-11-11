package service

import (
	"hackaton-no-code-constructor/pkg/repository"
)

type ProjectBlockService struct {
	repo repository.ProjectBlock
}

func NewProjectBlockService(repo repository.ProjectBlock) *ProjectBlockService {
	return &ProjectBlockService{repo: repo}
}

func (s *ProjectBlockService) Create() {}

func (s *ProjectBlockService) GetAll() {}

func (s *ProjectBlockService) GetByID() {}

func (s *ProjectBlockService) Update() {}

func (s *ProjectBlockService) Delete() {}
