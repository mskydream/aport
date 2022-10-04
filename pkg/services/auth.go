package services

import (
	"github.com/mskydream/aport/models"
	"github.com/mskydream/aport/pkg/repositories"
)

type AuthService struct {
	repo repositories.Auth
}

func NewAuthService(repo repositories.Auth) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) SignUp(user *models.UserProfile) (models.UserProfile, error) {
	return s.repo.SignUp(user)
}
