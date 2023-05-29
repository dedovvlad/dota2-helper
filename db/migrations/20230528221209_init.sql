-- +goose Up
-- +goose StatementBegin
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE heroes
(
    id        UUID             DEFAULT uuid_generate_v4() PRIMARY KEY,
    hero_name TEXT    NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
created_at
timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX ix_heroes_name ON heroes (hero_name);

CREATE TABLE items
(
    id        UUID             DEFAULT uuid_generate_v4() PRIMARY KEY,
    item_name TEXT    NOT NULL,
    link      TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
created_at
timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX ix_items_name ON items (item_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX ix_heroes_name;
DROP INDEX ix_items_name;
DROP TABLE heroes;
DROP TABLE items;
-- +goose StatementEnd
