package main

import (
	"fmt"
	"log"

	"github.com/Des331150/expense-tracker-api/internal/database"
)

func main() {
	pool, err := database.ConnectionPool()
	if err != nil {
		log.Fatalf("Fatal: %v\n", err)
	}

	defer pool.Close()

	fmt.Println("Database connection secured!")
}
