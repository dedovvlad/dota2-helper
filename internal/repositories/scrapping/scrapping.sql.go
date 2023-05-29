// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: scrapping.sql

package postgres

import (
	"context"
)

const newHeroes = `-- name: NewHeroes :exec
INSERT INTO heroes (hero_name)
VALUES ($1)
`

func (q *Queries) NewHeroes(ctx context.Context, heroName string) error {
	_, err := q.db.ExecContext(ctx, newHeroes, heroName)
	return err
}

const newItems = `-- name: NewItems :exec
INSERT INTO items (item_name)
VALUES ($1)
`

func (q *Queries) NewItems(ctx context.Context, itemName string) error {
	_, err := q.db.ExecContext(ctx, newItems, itemName)
	return err
}