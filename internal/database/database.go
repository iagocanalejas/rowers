package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Health() map[string]string

	// users.go
	GetUserById(userID int64) (*User, error)
	GetUsers() ([]User, error)
	CreateUser(user User) (*User, error)
	DeleteUser(userID int64) error

	// weights.go
	GetWeightsByUserId(userID int64) ([]Weight, error)
	AddWeight(userID int64, weight float64) error
	DeleteWeight(userID int64, weightId int64) error

	// assistances.go
	GetAssistanceById(assistanceID int64) (*Assistance, error)
	GetAssistances() ([]Assistance, error)
	CreateAssistance(assistance Assistance) (*Assistance, error)
	DeleteAssistance(assistanceID int64) error

	// user_assistances.go
	GetAssistanceByUserId(userID int64) ([]Assistance, error)
	AddUserAssistance(userID int64, assistanceID int64) error
	DeleteUserAssistance(userID int64, assistanceID int64) error
}

type service struct {
	db *sqlx.DB
}

func New() Service {
	conn, err := sqlx.Connect("sqlite3", os.Getenv("TURSO_DB_URL"))
	if err != nil {
		// This will not be a connection error, but a DSN parse error or another initialization error.
		log.Fatal(err)
	}
	s := &service{db: conn}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
