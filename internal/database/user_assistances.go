package database

import (
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (s *service) GetAssistanceByUserId(userID int64) ([]Assistance, error) {
	query, args, err := sq.
		Select("a.id", "a.type", "a.date").
		From("user_assistances ua").Join("assistances a ON ua.assistance_id = a.id").
		Where(sq.Eq{"user_id": userID}).
		OrderBy("date DESC").
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var assistances []Assistance
	if err = s.db.Select(&assistances, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return assistances, nil
}

func (s *service) AddUserAssistance(userID int64, assistanceID int64) error {
	query, args, err := sq.
		Insert("user_assistances").
		Columns("user_id", "assistance_id").
		Values(userID, assistanceID).
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

func (s *service) DeleteUserAssistance(userID int64, assistanceID int64) error {
	query, args, err := sq.
		Delete("user_assistances").
		Where(sq.Eq{"user_id": userID, "assistance_id": assistanceID}).
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
