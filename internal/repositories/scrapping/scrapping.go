package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Storage struct {
	db           *sqlx.DB
	q            *Queries
	queryBuilder squirrel.StatementBuilderType
}

func NewStorage(db *sqlx.DB) *Storage {
	return &Storage{
		db:           db,
		q:            New(db),
		queryBuilder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (s *Storage) AddHeroes(heroes []Hero) error {
	query := s.queryBuilder.Insert("heroes").
		Columns("hero_name")
	for _, hero := range heroes {
		query = query.Values(hero.HeroName)
	}

	stmt, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "making sql 'insert' from builder for 'heroes'")
	}

	_, err = s.db.Exec(stmt, args...)
	if err != nil {
		return errors.Wrap(err, "executing query for 'heroes'")
	}

	return nil
}

func (s *Storage) AddItems(items []Item) error {
	query := s.queryBuilder.Insert("items").
		Columns("item_name")
	for _, hero := range items {
		query = query.Values(hero.ItemName)
	}

	stmt, args, err := query.ToSql()
	if err != nil {
		return errors.Wrap(err, "making sql 'insert' from builder for 'items'")
	}

	_, err = s.db.Exec(stmt, args...)
	if err != nil {
		return errors.Wrap(err, "executing query for 'items'")
	}

	return nil
}
