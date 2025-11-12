package service

import (
	"fmt"
	"hackaton-no-code-constructor/pkg/dto/block_type_context"
	models "hackaton-no-code-constructor/pkg/model"
	"hackaton-no-code-constructor/pkg/repository"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BlockTypeService struct {
	repo repository.BlockType
}

func NewBlockTypeService(repo repository.BlockType) *BlockTypeService {
	return &BlockTypeService{repo: repo}
}

func (s *BlockTypeService) Create(input block_type_context.CreateBlockTypeInput) (*models.BlockType, error) {
	if !s.checkTag(input.TagID) {
		return nil, fmt.Errorf("tag %w", gorm.ErrRecordNotFound)
	}

	blockType := models.BlockType{
		TagID:       input.TagID,
		Name:        input.Name,
		Description: input.Description,
		Template:    input.Template,
		Schema:      input.Schema,
		Preview:     input.Preview,
	}

	createdBlockType, err := s.repo.Create(blockType)
	if err != nil {
		return nil, err
	}

	return createdBlockType, nil
}

func (s *BlockTypeService) GetAll() ([]models.BlockType, error) {
	blockTypes, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return blockTypes, nil
}

func (s *BlockTypeService) GetByID(id string) (*models.BlockType, error) {
	blockType, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return blockType, nil
}

func (s *BlockTypeService) Update(id string, input block_type_context.UpdateBlockTypeInput) (*models.BlockType, error) {
	if !s.checkTag(input.TagID) {
		return nil, fmt.Errorf("tag %w", gorm.ErrRecordNotFound)
	}

	blockType, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	blockType.TagID = input.TagID
	blockType.Name = input.Name
	blockType.Description = input.Description
	blockType.Template = input.Template
	blockType.Schema = input.Schema
	blockType.Preview = input.Preview

	updatedBlockType, err := s.repo.Update(blockType)
	if err != nil {
		return nil, err
	}

	return updatedBlockType, nil
}

func (s *BlockTypeService) Delete(id string) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s *BlockTypeService) checkTag(id uuid.UUID) bool {
	_, err := s.GetByID(id.String())
	return err == nil
}
