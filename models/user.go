package models

import "time"

type UserProfile struct {
	Id          string    `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Surname     string    `db:"surname" json:"surname"`
	Born        time.Time `db:"born" json:"born"`
	PhoneNumber string    `db:"phone_number" json:"phone_number"`
	Email       string    `db:"email" json:"email"`
	Gender      string    `db:"gender" json:"gender"`
	Address     string    `db:"address" json:"address"`
	Password    string    `db:"password" json:"password"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}
