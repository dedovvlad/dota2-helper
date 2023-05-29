-- name: NewHeroes :exec
INSERT INTO heroes (hero_name)
VALUES (@hero_name)
;

-- name: NewItems :exec
INSERT INTO items (item_name, link)
VALUES (@item_name, @link)
;

-- name: countHeroes :one
SELECT COUNT(*)
FROM heroes
;

-- name: countItems :one
SELECT COUNT(*)
FROM items
;

-- name: heroes :many
SELECT *
FROM heroes
;

-- name: items :many
SELECT *
FROM items
;