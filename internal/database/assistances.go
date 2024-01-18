package database

import (
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (s *service) GetAssistanceByID(assistanceID int64) (*Assistance, error) {
	query, args, err := sq.
		Select("id", "type", "date").
		From("assistances").
		Where(sq.Eq{"id": assistanceID}).
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var assistance Assistance
	if err = s.db.Get(&assistance, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return &assistance, nil
}

func (s *service) GetAssistances() ([]Assistance, error) {
	query, args, err := sq.
		Select("id", "type", "date").
		From("assistances").
		OrderBy("date").
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

func (s *service) CreateAssistance(assistance Assistance) (*Assistance, error) {
	query, args, err := sq.
		Insert("assistances").
		Columns("type", "date").
		Values(assistance.Type, assistance.Date).
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if _, err := s.db.Exec(query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	var assistanceID int64
	if err = s.db.Get(&assistanceID, "select last_insert_rowid()"); err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(assistanceID)

	return s.GetAssistanceByID(assistanceID)
}

func (s *service) DeleteAssistance(assistanceID int64) error {
	query, args, err := sq.
		Delete("assistances").
		Where(sq.Eq{"id": assistanceID}).
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
