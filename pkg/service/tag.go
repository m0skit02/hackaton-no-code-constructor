package service

import (
	"hackaton-no-code-constructor/pkg/dto/tag_context"
	models "hackaton-no-code-constructor/pkg/model"
	"hackaton-no-code-constructor/pkg/repository"
)

type TagService struct {
	repo repository.Tag
}

func NewTagService(repo repository.Tag) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) Create(input tag_context.CreateTagInput) (*models.Tag, error) {
	tag := models.Tag{
		Name: input.Name,
	}

	createdTag, err := s.repo.Create(tag)
	if err != nil {
		return nil, err
	}

	return createdTag, nil
}

func (s *TagService) GetAll() ([]models.Tag, error) {
	tags, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return tags, nil
}

func (s *TagService) GetByID(id string) (*models.Tag, error) {
	tag, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return tag, nil
}

func (s *TagService) Update(id string, input tag_context.UpdateTagInput) (*models.Tag, error) {
	tag, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	tag.Name = input.Name

	updatedTag, err := s.repo.Update(tag)
	if err != nil {
		return nil, err
	}

	return updatedTag, nil
}

func (s *TagService) Delete(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
