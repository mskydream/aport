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

func (r *Repo) SignUp(user *models.UserProfile) (int, error) {
	var id int
	query := `INSERT INTO user_profile (name, surname, born, phone_number, email, gender, address, password, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())`
	row := r.db.DB.QueryRow(query, user.Name, user.Surname, user.Born, user.PhoneNumber, user.Email, user.Gender, user.Address, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// tx, err := r.conn.Beginx()
// 	if err != nil {
// 		return
// 	}

// 	defer tx.Rollback()

// 	res, err := tx.PrepareNamed(`INSERT INTO user_profile (name, surname, born, phone_number, email, gender, address, password, created_at)
// 								VALUES :name, :surname, :born, :phone_number, :email, :gender, :address, :password, NOW()`)
// 	if err != nil {
// 		return
// 	}

// 	if err = res.QueryRowx(&user).Scan(&id); err != nil {
// 		return
// 	}

// 	return
