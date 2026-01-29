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

func (u *UserModel) CreateUser(email string, s string) error {
	panic("unimplemented")
}
