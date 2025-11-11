package service

import (
	"hackaton-no-code-constructor/pkg/repository"
)

type BlockTypeService struct {
	repo repository.BlockType
}

func NewBlockTypeService(repo repository.BlockType) *BlockTypeService {
	return &BlockTypeService{repo: repo}
}

func (s *BlockTypeService) Create() {}

func (s *BlockTypeService) GetAll() {}

func (s *BlockTypeService) GetByID() {}

func (s *BlockTypeService) Update() {}

func (s *BlockTypeService) Delete() {}
