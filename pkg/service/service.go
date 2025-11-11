package service

import (
	models "aeza-checker/pkg/model"
	"aeza-checker/pkg/repository"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Task interface {
	Create(task models.Task) (models.Task, error)
	GetAll() ([]models.Task, error)
	GetByID(id string) (models.Task, error)
	Update(task models.Task) (models.Task, error)
	Delete(id string) error
}

type Agent interface {
	Create(agent *models.Agent) (*models.Agent, error)
	GetByID(id uuid.UUID) (models.Agent, error)
	GetByToken(token string) (models.Agent, error)
	Update(agent *models.Agent) error
	Delete(id uuid.UUID) error
	List(status *string) ([]models.Agent, error)
	UpdateHeartbeat(id uuid.UUID) error
	SetStatus(id uuid.UUID, status string) error
}

type Result interface {
	SaveResult(result *models.Result) error
	GetResultsByTaskID(taskID uuid.UUID) ([]models.Result, error)
	GetLatestByAgent(agentID uuid.UUID, limit int) ([]models.Result, error)
	GetAll() ([]models.Result, error)
	Delete(id uuid.UUID) error
}

type Metrics interface {
	IncrementTodayChecks() error
	GetAll() ([]models.DailyMetric, error)
}

// Главная структура сервисов
type Service struct {
	Task
	Agent
	Result
	Metrics
}

// Конструктор
func NewService(repos *repository.Repository, rdb *redis.Client) *Service {
	return &Service{
		Task:    NewTaskService(repos.Task, rdb),
		Agent:   NewAgentService(repos.Agent),
		Result:  NewResultService(repos.Result, repos.Metrics),
		Metrics: NewMetricsService(repos.Metrics),
	}
}
