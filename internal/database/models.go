package database

import "time"

type User struct {
	ID        int64    `db:"id" json:"id"`
	FirstName string   `db:"first_name" json:"first_name"`
	LastName  string   `db:"last_name" json:"last_name"`
	Weight    *float64 `db:"weight" json:"weight"`
}

type Weight struct {
	ID     int64      `db:"id" json:"id"`
	UserID int64      `db:"user_id" json:"user_id"`
	Weight float64    `db:"weight" json:"weight"`
	Date   *time.Time `db:"date" json:"date"`
}
