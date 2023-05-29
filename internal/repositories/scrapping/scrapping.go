package postgres

import (
	"context"

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
		Columns("item_name", "link")
	for _, item := range items {
		query = query.Values(item.ItemName, item.Link)
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

func (s *Storage) CountHeroes(ctx context.Context) (int64, error) {
	count, err := s.q.countHeroes(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "selecting count heroes")
	}

	return count, nil
}

func (s *Storage) CountItems(ctx context.Context) (int64, error) {
	count, err := s.q.countItems(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "selecting count items")
	}

	return count, nil
}
