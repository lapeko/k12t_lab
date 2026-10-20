package store

import (
	"context"
	"errors"

	"github.com/lib/pq"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/model"
)

func (s *Store) ListBooks(ctx context.Context) ([]model.BookResponse, error) {
	query := `
        SELECT 
            b.id, 
            b.title,
            COALESCE(array_agg(a.name) FILTER (WHERE a.name IS NOT NULL), '{}') AS authors
        FROM books b
        LEFT JOIN LATERAL unnest(b.author_ids) AS aid ON true
        LEFT JOIN authors a ON a.id = aid
        GROUP BY b.id, b.title
        ORDER BY b.id
    `

	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []model.BookResponse{}
	for rows.Next() {
		var book model.BookResponse
		var authors pq.StringArray

		if err := rows.Scan(&book.ID, &book.Title, &authors); err != nil {
			return nil, err
		}

		book.Authors = []string(authors)
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (s *Store) CreateBook(ctx context.Context, title string, authorIDs []int64) (model.Book, error) {
	var authorsCount int
	err := s.db.QueryRowContext(
		ctx,
		`SELECT COUNT(*) FROM authors WHERE id = ANY($1)`,
		pq.Array(authorIDs),
	).Scan(&authorsCount)
	if err != nil {
		return model.Book{}, err
	}
	if authorsCount != len(uniqueIDs(authorIDs)) {
		return model.Book{}, ErrAuthorsMissing
	}

	book := model.Book{Title: title, AuthorIDs: authorIDs}
	err = s.db.QueryRowContext(
		ctx,
		`INSERT INTO books (title, author_ids) VALUES ($1, $2) RETURNING id`,
		book.Title,
		pq.Array(book.AuthorIDs),
	).Scan(&book.ID)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			return model.Book{}, ErrDuplicate
		}
		return model.Book{}, err
	}

	return book, nil
}
