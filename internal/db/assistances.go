package db

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type AssistanceType = string

const (
	AssistanceSea         AssistanceType = "SEA"
	AssistanceGym         AssistanceType = "GYM"
	AssistanceCompetition AssistanceType = "COMPETITION"
)

type Assistance struct {
	ID   int64          `db:"id" query:"id" param:"id" json:"id"`
	Type AssistanceType `db:"type" json:"type"`
	Date sql.NullTime   `db:"date" json:"date"`
}

func (r *Repository) GetAssistanceByID(assistanceID int64) (*Assistance, error) {
	query, args, err := sq.
		Select("id", "type", "date").
		From("assistances").
		Where(sq.Eq{"id": assistanceID}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var assistance Assistance
	if err = r.db.Get(&assistance, query, args...); err != nil {
		return nil, err
	}

	return &assistance, nil
}

func (r *Repository) GetAssistances() ([]Assistance, error) {
	query, args, err := sq.
		Select("id", "type", "date").
		From("assistances").
		OrderBy("date DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	var assistances []Assistance
	if err = r.db.Select(&assistances, query, args...); err != nil {
		return nil, err
	}

	return assistances, nil
}

func (r *Repository) CreateAssistance(a Assistance) (*Assistance, error) {
	query, args, err := sq.
		Insert("assistances").
		Columns("type", "date").
		Values(a.Type, a.Date.Time).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	var assistance Assistance
	if err = r.db.Get(&assistance, query, args...); err != nil {
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
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
