package service

import (
	"hackaton-no-code-constructor/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Create() {}

func (s *UserService) GetAll() {}

func (s *UserService) GetByID() {}

func (s *UserService) Update() {}

func (s *UserService) Delete() {}
