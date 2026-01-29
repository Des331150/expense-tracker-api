package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int64     `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	PasswordHash []byte    `json:"-"`
	CreatedAt    time.Time `json:"created_at"`
}

type UserModel struct {
	DB *sql.DB
}

func (u *UserModel) CreateUser(email string, PasswordHash string) error {
	query := `INSERT INTO users(email, password_hash) VALUES($1, $2)`
	_, err := u.DB.Exec(query, email, PasswordHash)
	return err
}
