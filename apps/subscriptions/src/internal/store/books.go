package store

import (
	"context"
	"errors"

	"github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/model"
)

func (s *Store) ListBooks(ctx context.Context) ([]model.BookResponse, error) {
	s.logger.Debug("fetching all books from database")

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
		s.logger.WithError(err).Error("failed to query books")
		return nil, err
	}
	defer rows.Close()

	books := []model.BookResponse{}
	for rows.Next() {
		var book model.BookResponse
		var authors pq.StringArray

		if err := rows.Scan(&book.ID, &book.Title, &authors); err != nil {
			s.logger.WithError(err).Error("failed to scan book row")
			return nil, err
		}

		book.Authors = []string(authors)
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		s.logger.WithError(err).Error("error occurred during rows iteration for books")
		return nil, err
	}

	s.logger.WithField("count", len(books)).Debug("successfully fetched books")
	return books, nil
}

func (s *Store) CreateBook(ctx context.Context, title string, authorIDs []int64) (model.Book, error) {
	s.logger.WithFields(logrus.Fields{
		"title":      title,
		"author_ids": authorIDs,
	}).Debug("checking authors existence")

	var authorsCount int
	err := s.db.QueryRowContext(
		ctx,
		`SELECT COUNT(*) FROM authors WHERE id = ANY($1)`,
		pq.Array(authorIDs),
	).Scan(&authorsCount)
	if err != nil {
		s.logger.WithError(err).Error("failed to check authors existence")
		return model.Book{}, err
	}

	if authorsCount != len(uniqueIDs(authorIDs)) {
		s.logger.Warn("one or more authors do not exist in database")
		return model.Book{}, ErrAuthorsMissing
	}

	book := model.Book{Title: title, AuthorIDs: authorIDs}
	s.logger.WithField("title", book.Title).Debug("inserting new book into database")

	err = s.db.QueryRowContext(
		ctx,
		`INSERT INTO books (title, author_ids) VALUES ($1, $2) RETURNING id`,
		book.Title,
		pq.Array(book.AuthorIDs),
	).Scan(&book.ID)

	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == "23505" {
			s.logger.WithError(err).Warn("attempt to create a duplicate book")
			return model.Book{}, ErrDuplicate
		}
		s.logger.WithError(err).Error("failed to insert book")
		return model.Book{}, err
	}

	s.logger.WithField("book_id", book.ID).Info("book successfully created")
	return book, nil
}
