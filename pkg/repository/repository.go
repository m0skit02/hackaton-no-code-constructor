package repository

import (
	"github.com/google/uuid"
	models "hackaton-no-code-constructor/pkg/model"

	"gorm.io/gorm"
)

type User interface {
	Create(user models.User) (*models.User, error)
	GetByID(id uuid.UUID) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(id uuid.UUID) error
	GetAll() ([]models.User, error)
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
	Create(project models.Project) (*models.Project, error)
	GetAll() ([]models.Project, error)
	GetByID(id uuid.UUID) (*models.Project, error)
	Update(project models.Project) (*models.Project, error)
	Delete(id uuid.UUID) error
}

type ProjectBlock interface {
	Create(projectBlock models.ProjectBlock) (*models.ProjectBlock, error)
	GetAll() ([]models.ProjectBlock, error)
	GetByID(id uuid.UUID) (*models.ProjectBlock, error)
	Update(projectBlock models.ProjectBlock) (*models.ProjectBlock, error)
	Delete(id uuid.UUID) error
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
