package database

import (
	"log"

	sq "github.com/Masterminds/squirrel"
)

type User struct {
	Id        int64    `db:"id" json:"id"`
	FirstName string   `db:"first_name" json:"first_name" form:"first_name"`
	LastName  string   `db:"last_name" json:"last_name" form:"last_name"`
	Weight    *float64 `db:"weight" json:"weight"`
}

func (s *service) GetUserById(userId int64) (*User, error) {
	query, args, err := sq.
		Select("id, first_name, last_name").
		From("users u").
		Where(sq.Eq{"id": userId}).
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var user User
	if err = s.db.Get(&user, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return &user, nil
}

func (s *service) GetUsers() ([]User, error) {
	query, args, err := sq.
		Select(
			"id", "first_name", "last_name",
			"(SELECT LAST_VALUE(weight) OVER (ORDER BY creation_date DESC RANGE BETWEEN UNBOUNDED PRECEDING AND UNBOUNDED FOLLOWING) FROM weights w WHERE w.user_id = u.id ORDER BY creation_date DESC LIMIT 1) as weight",
		).
		From("users u").
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var users []User
	if err = s.db.Select(&users, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return users, nil
}

func (s *service) CreateUser(u User) (*User, error) {
	query, args, err := sq.
		Insert("users").
		Columns("first_name", "last_name").
		Values(u.FirstName, u.LastName).
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if _, err := s.db.Exec(query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	var userId int64
	if err = s.db.Get(&userId, "select last_insert_rowid()"); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(userId)

	return s.GetUserById(userId)
}

func (s *service) DeleteUser(userId int64) error {
	query, args, err := sq.
		Delete("users").
		Where(sq.Eq{"id": userId}).
		ToSql()
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err := s.db.Exec(query, args...); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
