-- +goose Up
CREATE TABLE IF NOT EXISTS authors (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS books (
    id BIGSERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author_ids BIGINT [] NOT NULL,
    CONSTRAINT unique_title_authors UNIQUE (title, author_ids)
);

CREATE TABLE IF NOT EXISTS inventory (
    book_id BIGINT PRIMARY KEY REFERENCES books(id) ON DELETE CASCADE,
    quantity INT NOT NULL DEFAULT 0 CHECK (quantity >= 0)
);

-- +goose Down
DROP TABLE IF EXISTS inventory;
DROP TABLE IF EXISTS books;
DROP TABLE IF EXISTS authors;