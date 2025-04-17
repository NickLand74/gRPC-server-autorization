package storage

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/NickLand74/gRPC-server-autorization/config"
)

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() *PostgresStorage {
	cfg := config.LoadConfig()
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))
	if err != nil {
		log.Fatal(err)
	}
	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) CreateUser(username, password string) error {
	_, err := s.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	return err
}

func (s *PostgresStorage) GetUser(username string) (*User, error) {
	var user User
	err := s.db.QueryRow("SELECT username, password FROM users WHERE username = $1", username).Scan(&user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, err
	}
	return &user, err
}
