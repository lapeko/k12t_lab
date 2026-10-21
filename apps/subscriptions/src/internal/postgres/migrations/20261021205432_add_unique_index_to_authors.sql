-- +goose Up
CREATE UNIQUE INDEX authors_name_idx ON authors (name);

-- +goose Down
DROP INDEX IF EXISTS authors_name_idx;