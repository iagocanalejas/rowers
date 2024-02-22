package db

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type User struct {
	ID              int64           `db:"id" json:"id"`
	FirstName       string          `db:"first_name" json:"first_name"`
	LastName        string          `db:"last_name" json:"last_name"`
	Weight          sql.NullFloat64 `db:"weight" json:"weight"`
	Assistance      sql.NullFloat64 `db:"assistance" json:"assistance"`
	TotalAssistance sql.NullFloat64 `db:"total_assistance" json:"total_assistance"`
}

func (r *Repository) GetUserByID(userId int64) (*User, error) {
	query, args, err := sq.
		Select("id", "first_name", "last_name").
		From("users u").
		Where(sq.Eq{"id": userId}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var user User
	if err = r.db.Get(&user, query, args...); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) GetUsers() ([]User, error) {
	now := time.Now()
	firstDayOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	weightQuery := "(SELECT LAST_VALUE(weight) OVER (ORDER BY date ASC RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING) FROM weights w WHERE w.user_id = u.id) as weight"
	assistanceQuery := "(SELECT COUNT(*) FROM assistances WHERE date >= ?) as total_assistance"
	userAssistanceQuery := "(SELECT COUNT(*) FROM user_assistances ua JOIN assistances a ON ua.assistance_id = a.id WHERE ua.user_id = u.id AND a.date >= ?) as assistance"

	query, args, err := sq.
		Select("id", "first_name", "last_name", weightQuery, assistanceQuery, userAssistanceQuery).
		From("users u").
		OrderBy("first_name", "last_name").
		ToSql()
	if err != nil {
		return nil, err
	}

	args = append(args, firstDayOfMonth.Format("2006-01-02"), firstDayOfMonth.Format("2006-01-02"))

	var users []User
	if err = r.db.Select(&users, query, args...); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) CreateUser(u User) (*User, error) {
	query, args, err := sq.
		Insert("users").
		Columns("first_name", "last_name").
		Values(u.FirstName, u.LastName).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	var user User
	if err = r.db.Get(&user, query, args...); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) DeleteUser(userId int64) error {
	query, args, err := sq.
		Delete("users").
		Where(sq.Eq{"id": userId}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
