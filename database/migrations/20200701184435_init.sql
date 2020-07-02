-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE IF NOT EXISTS items (
  id          INTEGER PRIMARY KEY,
  name        TEXT NOT NULL UNIQUE,
  img_url     TEXT,
  stock       INTEGER default 0,
  unit        TEXT NOT NULL,
  min_stock   INTEGER default 0
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE IF EXISTS items;
