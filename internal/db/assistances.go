package db

import (
	"log"

	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) GetAssistanceById(assistanceID int64) (*Assistance, error) {
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
	if err = r.db.Get(&assistance, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return &assistance, nil
}

func (r *Repository) GetAssistances() ([]Assistance, error) {
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
	if err = r.db.Select(&assistances, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return assistances, nil
}

func (r *Repository) CreateAssistance(a Assistance) (*Assistance, error) {
	query, args, err := sq.
		Insert("assistances").
		Columns("type", "date").
		Values(a.Type, a.Date).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var assistance Assistance
	if err = r.db.Get(&assistance, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return &assistance, nil
}

func (r *Repository) DeleteAssistance(assistanceID int64) error {
	query, args, err := sq.
		Delete("assistances").
		Where(sq.Eq{"id": assistanceID}).
		ToSql()
	if err != nil {
		log.Println(err)
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
