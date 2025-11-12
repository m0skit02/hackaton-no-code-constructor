package service

import (
	"hackaton-no-code-constructor/pkg/dto/block_type_context"
	models "hackaton-no-code-constructor/pkg/model"
	"hackaton-no-code-constructor/pkg/repository"
)

type User interface {
	Create()
	GetAll()
	GetByID()
	Update()
	Delete()
}

type Tag interface {
	Create()
	GetAll()
	GetByID()
	Update()
	Delete()
}

type BlockType interface {
	Create(input block_type_context.CreateBlockTypeInput) (*models.BlockType, error)
	GetAll() ([]models.BlockType, error)
	GetByID(id string) (*models.BlockType, error)
	Update(id string, input block_type_context.UpdateBlockTypeInput) (*models.BlockType, error)
	Delete(id string) error
}

type Project interface {
	Create()
	GetAll()
	GetByID()
	Update()
	Delete()
}

type ProjectBlock interface {
	Create()
	GetAll()
	GetByID()
	Update()
	Delete()
}

// Главная структура сервисов
type Service struct {
	User
	Tag
	BlockType
	Project
	ProjectBlock
}

// Конструктор
func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:         NewUserService(repos.User),
		Tag:          NewTagService(repos.Tag),
		BlockType:    NewBlockTypeService(repos.BlockType),
		Project:      NewProjectService(repos.Project),
		ProjectBlock: NewProjectBlockService(repos.ProjectBlock),
	}
}
