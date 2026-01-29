package models

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
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
	DB *pgxpool.Pool
}

func (u *UserModel) CreateUser(email string, PasswordHash string) error {
	query := `INSERT INTO users(email, password_hash) VALUES($1, $2)`
	_, err := u.DB.Exec(query, email, PasswordHash)
	return err
}

func (u *UserModel) LogUser(email string) (string, int, error) {
	var hash string
	var id int
	query := `SELECT password_hash FROM users WHERE email = $1`
	err := u.DB.QueryRow(context.Background(), query, email).Scan(&hash, &id)
	if err != nil {
		return "", 0, err
	}

	return hash, id, nil
}
