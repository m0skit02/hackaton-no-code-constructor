package service

import (
	"hackaton-no-code-constructor/pkg/repository"
)

type TagService struct {
	repo repository.Tag
}

func NewTagService(repo repository.Tag) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) Create() {}

func (s *TagService) GetAll() {}

func (s *TagService) GetByID() {}

func (s *TagService) Update() {}

func (s *TagService) Delete() {}
