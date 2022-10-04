package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/aport/models"
)

type Auth interface {
	SignUp(user *models.UserProfile) (models.UserProfile, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepository(db),
	}
}
