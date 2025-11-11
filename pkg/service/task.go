package service

import (
	models "aeza-checker/pkg/model"
	"aeza-checker/pkg/repository"
	"context"
	"encoding/json"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type TaskService struct {
	repo  repository.Task
	redis *redis.Client
}

func NewTaskService(repo repository.Task, rdb *redis.Client) *TaskService {
	return &TaskService{repo: repo, redis: rdb}
}

func (s *TaskService) Create(task models.Task) (models.Task, error) {
	created, err := s.repo.Create(task)
	if err != nil {
		return models.Task{}, err
	}

	// Добавляем задачу в Redis очередь
	data, _ := json.Marshal(created)
	err = s.redis.LPush(context.Background(), "task_queue", data).Err()
	if err != nil {
		fmt.Println("⚠️ Redis push error:", err)
	}

	return created, nil
}

func (s *TaskService) GetAll() ([]models.Task, error) {
	return s.repo.GetAll()
}

func (s *TaskService) GetByID(id string) (models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *TaskService) Update(task models.Task) (models.Task, error) {
	return s.repo.Update(task)
}

func (s *TaskService) Delete(id string) error {
	return s.repo.Delete(id)
}
