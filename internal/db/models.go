package db

import (
	"database/sql"
)

type User struct {
	ID              int64           `db:"id" json:"id"`
	FirstName       string          `db:"first_name" json:"first_name"`
	LastName        string          `db:"last_name" json:"last_name"`
	Weight          sql.NullFloat64 `db:"weight" json:"weight"`
	Assistance      sql.NullFloat64 `db:"assistance" json:"assistance"`
	TotalAssistance sql.NullFloat64 `db:"total_assistance" json:"total_assistance"`
}

type Weight struct {
	ID     int64        `db:"id" json:"id"`
	UserID int64        `db:"user_id" json:"user_id"`
	Weight float64      `db:"weight" json:"weight"`
	Date   sql.NullTime `db:"date" json:"date"`
}

type AssistanceType = string

const (
	AssistanceSea         AssistanceType = "SEA"
	AssistanceGym         AssistanceType = "GYM"
	AssistanceCompetition AssistanceType = "COMPETITION"
)

type Assistance struct {
	ID   int64          `db:"id" json:"id"`
	Type AssistanceType `db:"type" json:"type"`
	Date sql.NullTime   `db:"date" json:"date"`
}

type UserAssistance struct {
	UserID       *int64         `db:"user_id" json:"user_id"`
	AssistanceID int64          `db:"assistance_id" json:"assistance_id"`
	Type         AssistanceType `db:"type" json:"type"`
	Date         sql.NullTime   `db:"date" json:"date"`
}
