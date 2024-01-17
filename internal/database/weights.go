package database

import (
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (s *service) GetWeightsByUserId(userID int64) ([]Weight, error) {
	query, args, err := sq.
		Select("id", "user_id", "weight", "creation_date").
		From("weights").
		Where(sq.Eq{"user_id": userID}).
		OrderBy("creation_date DESC").
		ToSql()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var weights []Weight
	if err = s.db.Select(&weights, query, args...); err != nil {
		log.Println(err)
		return nil, err
	}

	return weights, nil
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

func (s *service) DeleteWeight(userId int64, weightId int64) error {
	query, args, err := sq.
		Delete("weights").
		Where(sq.Eq{"id": weightId, "user_id": userId}).
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
