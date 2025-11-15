package service

import (
	"github.com/google/uuid"
	"hackaton-no-code-constructor/pkg/dto/block_type_context"
	"hackaton-no-code-constructor/pkg/dto/project_block_context"
	"hackaton-no-code-constructor/pkg/dto/project_context"
	"hackaton-no-code-constructor/pkg/dto/tag_context"
	user_context "hackaton-no-code-constructor/pkg/dto/user_context"
	models "hackaton-no-code-constructor/pkg/model"
	"hackaton-no-code-constructor/pkg/repository"
)

type User interface {
	CreateUser(input user_context.CreateUserInput) (*models.User, error)
	LoginUser(username, password string) (*models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
	UpdateUser(id uuid.UUID, input user_context.UpdateUserInput) (*models.User, error)
	DeleteUser(id uuid.UUID) error
}

type Tag interface {
	Create(input tag_context.CreateTagInput) (*models.Tag, error)
	GetAll() ([]models.Tag, error)
	GetByIDTag(id string) (*models.Tag, error)
	Update(id string, input tag_context.UpdateTagInput) (*models.Tag, error)
	Delete(id string) error
}

type BlockType interface {
	Create(input block_type_context.CreateBlockTypeInput) (*models.BlockType, error)
	GetAll() ([]models.BlockType, error)
	GetByID(id string) (*models.BlockType, error)
	Update(id string, input block_type_context.UpdateBlockTypeInput) (*models.BlockType, error)
	Delete(id string) error
}

type Project interface {
	CreateProject(input project_context.CreateProjectInput) (*models.Project, error)
	GetAllProjects() ([]models.Project, error)
	GetByIDProject(id uuid.UUID) (*models.Project, error)
	UpdateProject(id uuid.UUID, input project_context.UpdateProjectInput) (*models.Project, error)
	DeleteProject(id uuid.UUID) error
	GetByUserID(userID string) ([]models.Project, error)
}

type ProjectBlock interface {
	CreateProjectBlock(input project_block_context.CreateProjectBlockInput) (*models.ProjectBlock, error)
	GetAllProjectBlock() ([]models.ProjectBlock, error)
	GetByIDProjectBlock(id uuid.UUID) (*models.ProjectBlock, error)
	UpdateProjectBlock(id uuid.UUID, input project_block_context.UpdateProjectBlockInput) (*models.ProjectBlock, error)
	DeleteProjectBlock(id uuid.UUID) error
}

type Auth interface {
	Login(username, password string) (*models.User, error)
}

// Главная структура сервисов
type Service struct {
	User
	Tag
	BlockType
	Project
	ProjectBlock
	Auth
}

// Конструктор
func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:         NewUserService(repos.User),
		Tag:          NewTagService(repos.Tag),
		BlockType:    NewBlockTypeService(repos.BlockType, repos.Tag), // <- вот тут
		Project:      NewProjectService(repos.Project),
		ProjectBlock: NewProjectBlockService(repos.ProjectBlock),
		Auth:         NewAuthService(repos.Auth),
	}
}
