package db

import (
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type UserFault struct {
	UserID       sql.NullInt64 `db:"user_id" query:"user_id" param:"user_id" json:"user_id"`
	AssistanceID int64         `db:"assistance_id" query:"assistance_id" param:"assistance_id" json:"assistance_id"`
	IsJustified  bool          `db:"is_justified" json:"is_justified"`
	Cause        string        `db:"cause" json:"cause"`
}

// TODO: services for this
// TODO: return fault when retrieving user assistances
// TODO: how do we view this in the web?

func (r *Repository) CreateUserFault(f *UserFault) (*UserFault, error) {
	query, args, err := sq.
		Insert("user_faults").
		Columns("user_id", "assistance_id", "is_justified", "cause", "creation_date").
		Values(f.UserID.Int64, f.AssistanceID, f.IsJustified, f.Cause, time.Now()).
		Suffix("RETURNING *").
		ToSql()
	if err != nil {
		return nil, err
	}

	var fault UserFault
	if err = r.db.Get(&fault, query, args...); err != nil {
		return nil, err
	}

	return &fault, nil
}

func (r *Repository) DeleteUserFault(userID int64, assistanceID int64) error {
	query, args, err := sq.
		Delete("user_faults").
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
