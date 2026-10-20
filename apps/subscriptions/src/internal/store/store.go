package store

import (
	"database/sql"
	"errors"

	"github.com/sirupsen/logrus"
)

var (
	ErrDuplicate      = errors.New("duplicate")
	ErrAuthorsMissing = errors.New("authors missing")
)

type Store struct {
	db     *sql.DB
	logger *logrus.Entry
}

func New(db *sql.DB, log *logrus.Logger) *Store {
	return &Store{db: db, logger: log.WithField("component", "store")}
}

func uniqueIDs(ids []int64) []int64 {
	seen := make(map[int64]struct{}, len(ids))
	out := make([]int64, 0, len(ids))
	for _, id := range ids {
		if _, ok := seen[id]; ok {
			continue
		}
		seen[id] = struct{}{}
		out = append(out, id)
	}
	return out
}
