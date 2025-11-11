package repository

import (
	"gorm.io/gorm"
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
	Create()
	GetAll()
	GetByID()
	Update()
	Delete()
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
type Repository struct {
	User
	Tag
	BlockType
	Project
	ProjectBlock
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:         NewUserRepo(db),
		Tag:          NewTagRepo(db),
		BlockType:    NewBlockTypeRepo(db),
		Project:      NewProjectRepo(db),
		ProjectBlock: NewProjectBlockRepo(db),
	}
}
