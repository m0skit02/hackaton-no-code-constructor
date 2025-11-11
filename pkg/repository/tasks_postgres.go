package repository

import (
	models "aeza-checker/pkg/model"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewTaskRepo(db *gorm.DB, rdb *redis.Client) *TaskRepo {
	return &TaskRepo{db: db, rdb: rdb}
}

func (r *TaskRepo) Create(task models.Task) (models.Task, error) {
	if err := r.db.Create(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	if err := r.db.Order("created_at DESC").Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *TaskRepo) GetByID(id string) (models.Task, error) {
	var task models.Task
	if err := r.db.First(&task, "id = ?", id).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) Update(task models.Task) (models.Task, error) {
	if err := r.db.Save(&task).Error; err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepo) Delete(id string) error {
	return r.db.Delete(&models.Task{}, "id = ?", id).Error
}
