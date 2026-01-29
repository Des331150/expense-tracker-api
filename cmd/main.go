package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Des331150/expense-tracker-api/internal/database"
	"github.com/Des331150/expense-tracker-api/internal/models"
	"github.com/go-playground/validator/v10"
)

func main() {
	pool, err := database.ConnectionPool()
	if err != nil {
		log.Fatalf("Fatal: %v\n", err)
	}

	defer pool.Close()

	fmt.Println("Database connection secured!")

	v := validator.New(validator.WithRequiredStructEnabled())
	secret := os.Getenv("secret")
	userStore := &models.UserModel{DB: pool}

	s := &Server{
		store:     userStore,
		validate:  v,
		jwtSecret: []byte(secret),
	}
	_ = s
}
