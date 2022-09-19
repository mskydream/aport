package services

import (
	"github.com/mskydream/aport/models"
	"github.com/mskydream/aport/pkg/repositories"
)

type Auth interface {
	SignUp(user *models.UserProfile) (int, error)
}

type Service struct {
	Auth
}

func NewService(repos *repositories.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}
