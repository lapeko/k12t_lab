package store

import (
	"context"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/model"
)

func (s *Store) ListAuthors(ctx context.Context) ([]model.Author, error) {
	rows, err := s.db.QueryContext(ctx, `SELECT id, name FROM authors ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	authors := make([]model.Author, 0)
	for rows.Next() {
		var author model.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, author)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authors, nil
}

func (s *Store) CreateAuthor(ctx context.Context, name string) (model.Author, error) {
	author := model.Author{Name: name}
	err := s.db.QueryRowContext(
		ctx,
		`INSERT INTO authors (name) VALUES ($1) RETURNING id`,
		author.Name,
	).Scan(&author.ID)
	return author, err
}
