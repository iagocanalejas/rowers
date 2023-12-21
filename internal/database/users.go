package database

import (
	"log"

	sq "github.com/Masterminds/squirrel"
)

type User struct {
	Id        int64  `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

func (s *service) GetUsers() ([]User, error) {
	query, args, err := sq.
		Select("id, first_name, last_name").
		From("users u").
		ToSql()
	if err != nil {
		log.Print(query)
		return nil, err
	}

	var users []User
	err = s.db.Select(&users, query, args...)
	if err != nil {
		return nil, err
	}

	return users, nil
}
