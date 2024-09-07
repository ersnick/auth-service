package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
	}
	fmt.Println("Connected to the database!")
}

func SaveRefreshToken(userId string, refreshToken string) error {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO user_tokens (user_id, refresh_token) VALUES ($1, $2)`
	_, err = db.Exec(query, userId, hashedToken)
	return err
}

func ValidateRefreshToken(userId string, refreshToken string) (bool, error) {
	var storedHash string
	query := `SELECT refresh_token FROM user_tokens WHERE user_id = $1`
	err := db.QueryRow(query, userId).Scan(&storedHash)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(refreshToken))
	if err != nil {
		return false, err
	}
	return true, nil
}
