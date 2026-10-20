package store

import (
	"context"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/model"
	"github.com/lib/pq"
)

type Author interface {
	ListAuthors(ctx context.Context) ([]model.Author, error)
	CreateAuthors(ctx context.Context, names []string) ([]model.Author, error)
}

func (s *Store) ListAuthors(ctx context.Context) ([]model.Author, error) {
	s.logger.Debug("query authors")
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

	s.logger.Debug("query authors done")
	return authors, nil
}

func (s *Store) CreateAuthors(ctx context.Context, names []string) ([]model.Author, error) {
	if len(names) == 0 {
		return []model.Author{}, nil
	}

	query := `
		INSERT INTO authors (name) 
		SELECT unnest($1::text[]) 
		RETURNING id, name
	`

	rows, err := s.db.QueryContext(ctx, query, pq.Array(names))
	if err != nil {
		s.logger.WithError(err).Error("failed to batch create authors")
		return nil, err
	}
	defer rows.Close()

	var authors []model.Author
	for rows.Next() {
		var author model.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			s.logger.WithError(err).Error("failed to scan author row")
			return nil, err
		}
		authors = append(authors, author)
	}

	if err := rows.Err(); err != nil {
		s.logger.WithError(err).Error("error occurred during rows iteration")
		return nil, err
	}

	s.logger.WithField("count", len(authors)).Info("successfully batch created authors")
	return authors, nil
}
