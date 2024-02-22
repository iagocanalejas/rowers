package db

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
)

type UserAssistance struct {
	UserID       *int64         `db:"user_id" json:"user_id"`
	AssistanceID int64          `db:"assistance_id" json:"assistance_id"`
	Type         AssistanceType `db:"type" json:"type"`
	Date         sql.NullTime   `db:"date" json:"date"`
}

func (r *Repository) GetUserAssistanceByUserIDAndAssistanceID(userID int64, assistanceID int64) (*UserAssistance, error) {
	query, args, err := sq.
		Select("a.id as assistance_id", "a.type", "a.date", "ua.user_id").
		From("user_assistances ua").Join("assistances a ON ua.assistance_id = a.id").
		Where(sq.And{
			sq.Eq{"ua.user_id": userID},
			sq.Eq{"ua.assistance_id": assistanceID},
		}).
		ToSql()
	if err != nil {
		return nil, err
	}

	var assistance UserAssistance
	if err = r.db.Get(&assistance, query, args...); err != nil {
		return nil, err
	}

	return &assistance, nil
}

func (r *Repository) GetUserAssistancesByUserID(userID int64) ([]UserAssistance, error) {
	query, args, err := sq.
		Select("a.id as assistance_id", "a.type", "a.date", "ua.user_id").
		From("assistances a").LeftJoin("user_assistances ua ON ua.assistance_id = a.id").
		Where(sq.Or{
			sq.Eq{"user_id": userID},
			sq.Eq{"user_id": nil},
		}).
		OrderBy("date DESC").
		ToSql()
	if err != nil {
		return nil, err
	}

	var assistances []UserAssistance
	if err = r.db.Select(&assistances, query, args...); err != nil {
		return nil, err
	}

	return assistances, nil
}

func (r *Repository) CreateUserAssistance(userID int64, assistanceID int64) (*UserAssistance, error) {
	query, args, err := sq.
		Insert("user_assistances").
		Columns("user_id", "assistance_id").
		Values(userID, assistanceID).
		ToSql()
	if err != nil {
		return nil, err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return nil, err
	}

	return r.GetUserAssistanceByUserIDAndAssistanceID(userID, assistanceID)
}

func (r *Repository) DeleteUserAssistance(userID int64, assistanceID int64) error {
	query, args, err := sq.
		Delete("user_assistances").
		Where(sq.Eq{"user_id": userID, "assistance_id": assistanceID}).
		ToSql()
	if err != nil {
		return err
	}

	if _, err := r.db.Exec(query, args...); err != nil {
		return err
	}

	return nil
}
