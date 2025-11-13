package repository

import (
	models "hackaton-no-code-constructor/pkg/model"

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
	Create(models.Tag) (*models.Tag, error)
	GetAll() ([]models.Tag, error)
	GetByID(id string) (*models.Tag, error)
	Update(tag *models.Tag) (*models.Tag, error)
	Delete(id string) error
}

type BlockType interface {
	Create(blockType models.BlockType) (*models.BlockType, error)
	GetAll() ([]models.BlockType, error)
	GetByID(id string) (*models.BlockType, error)
	Update(blockType *models.BlockType) (*models.BlockType, error)
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

type Auth interface {
	GetUserByUsernameAndPasswordHash(username string, passwordHash string) (*models.User, error)
}
type Repository struct {
	User
	Tag
	BlockType
	Project
	ProjectBlock
	Auth
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		User:         NewUserRepo(db),
		Tag:          NewTagRepo(db),
		BlockType:    NewBlockTypeRepo(db),
		Project:      NewProjectRepo(db),
		ProjectBlock: NewProjectBlockRepo(db),
		Auth:         NewAuthRepo(db),
	}
}
