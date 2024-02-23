package db

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type Weight struct {
	ID     int64        `db:"id" query:"id" param:"id" json:"id"`
	UserID int64        `db:"user_id" query:"user_id" param:"user_id" json:"user_id"`
	Weight float64      `db:"weight" json:"weight,string"`
	Date   sql.NullTime `db:"date" json:"date"`
}

func (r *Repository) GetWeightsByUserID(userID int64) ([]Weight, error) {
	query, args, err := sq.
		Select("id", "user_id", "weight", "date").
		From("weights").
		Where(sq.Eq{"user_id": userID}).
		OrderBy("date DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	var weights []Weight
	if err = r.db.Select(&weights, query, args...); err != nil {
		return nil, err
	}

	return weights, nil
}

func (r *Repository) CreateWeight(userID int64, weightValue float64) (*Weight, error) {
	query, args, err := sq.
		Insert("weights").
		Columns("user_id", "weight", "date").
		Values(userID, weightValue, time.Now()).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	var weight Weight
	if err = r.db.Get(&weight, query, args...); err != nil {
		return nil, err
	}

	return &weight, nil
}

func (r *Repository) DeleteWeight(userID int64, weightId int64) error {
	query, args, err := sq.
		Delete("weights").
		Where(sq.Eq{"id": weightId, "user_id": userID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
