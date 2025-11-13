package service

import (
	models "hackaton-no-code-constructor/pkg/model"
	"hackaton-no-code-constructor/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (r *AuthService) GetUserByUsernameAndPasswordHash(username string, passwordHash string) (*models.User, error) {
	user, err := r.repo.GetUserByUsernameAndPasswordHash(username, passwordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}
