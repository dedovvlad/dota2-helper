-- name: NewHeroes :exec
INSERT INTO heroes (hero_name)
VALUES (@hero_name)
;

-- name: NewItems :exec
INSERT INTO items (item_name)
VALUES (@item_name)
;
