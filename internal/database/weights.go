package database

import (
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
)

type UserWeight struct {
	UserId       int64      `db:"user_id" json:"user_id"`
	Weight       float64    `db:"weight" json:"weight"`
	CreationDate *time.Time `db:"creation_date" json:"creation_date"`
}

func (s *service) AddWeight(userId int64, weight float64) error {
	query, args, err := sq.
		Insert("weights").
		Columns("user_id", "weight", "creation_date").
		Values(userId, weight, time.Now()).
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

func (s *service) GetWeights(userId int64) ([]UserWeight, error) {
	query, args, err := sq.
		Select("user_id", "weight", "creation_date").
		From("weights").
		Where(sq.Eq{"user_id": userId}).
		OrderBy("creation_date DESC").
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var weights []UserWeight
	if err = s.db.Select(&weights, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return weights, nil
}
