package repository

import (
	models "hackaton-no-code-constructor/pkg/model"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Task interface {
	Create(task models.Task) (models.Task, error)
	GetAll() ([]models.Task, error)
	GetByID(id string) (models.Task, error)
	Update(task models.Task) (models.Task, error)
	Delete(id string) error
}

type Result interface {
	SaveResult(result *models.Result) error
	GetResultsByTaskID(taskID uuid.UUID) ([]models.Result, error)
	GetLatestByAgent(agentID uuid.UUID, limit int) ([]models.Result, error)
	GetAll() ([]models.Result, error)
	Delete(id uuid.UUID) error
}

type Agent interface {
	Create(agent *models.Agent) (uuid.UUID, error)
	GetByID(id uuid.UUID) (models.Agent, error)
	GetByToken(token string) (models.Agent, error)
	Update(agent *models.Agent) error
	Delete(id uuid.UUID) error
	List(status *string) ([]models.Agent, error)
	UpdateHeartbeat(id uuid.UUID, t time.Time) error
	SetStatus(id uuid.UUID, status string) error
}

type Metrics interface {
	IncrementTodayChecks() error
	GetAll() ([]models.DailyMetric, error)
}
type Repository struct {
	Task
	Result
	Agent
	Metrics
}

func NewRepository(db *gorm.DB, rdb *redis.Client) *Repository {
	return &Repository{
		Task:    NewTaskRepo(db, rdb),
		Result:  NewResultRepo(db, rdb),
		Agent:   NewAgentRepo(db, rdb),
		Metrics: NewMetricRepo(db),
	}
}
