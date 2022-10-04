package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/mskydream/aport/models"
)

type Repo struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *Repo {
	return &Repo{db: db}
}

func (r *Repo) SignUp(user *models.UserProfile) (model models.UserProfile, err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return
	}

	defer tx.Rollback()

	query := `INSERT INTO user_profile (name, surname, born, phone_number, email, gender, address, password, created_at) 
				VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW()) 
				RETURNING id, name, surname, born, phone_number, email, gender, address, password, created_at`
	row := r.db.DB.QueryRow(query, user.Name, user.Surname, user.Born, user.PhoneNumber, user.Email, user.Gender, user.Address, user.Password)

	if err = row.Scan(&model.Id, &model.Name, &model.Surname, &model.Born, &model.PhoneNumber, &model.Email, &model.Gender, &model.Address, &model.Password, &model.CreatedAt); err != nil {
		return
	}

	return
}
