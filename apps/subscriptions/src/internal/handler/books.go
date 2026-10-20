package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/model"
	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/store"
	"github.com/sirupsen/logrus"
)

type Books struct {
	store  store.Book
	logger *logrus.Entry
}

func NewBooks(store store.Book, log *logrus.Logger) *Books {
	return &Books{
		store:  store,
		logger: log.WithField("component", "books_handler"),
	}
}

func (h *Books) List(w http.ResponseWriter, r *http.Request) {
	h.logger.Debug("handling list books request")

	books, err := h.store.ListBooks(r.Context())
	if err != nil {
		h.logger.WithError(err).Error("failed to list books from store")
		writeError(w, http.StatusInternalServerError, "failed to list books")
		return
	}

	h.logger.WithField("count", len(books)).Debug("successfully listed books")
	writeJSON(w, http.StatusOK, books)
}

func (h *Books) Create(w http.ResponseWriter, r *http.Request) {
	var req model.Book
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.WithError(err).Warn("failed to decode json for book creation")
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if req.Title == "" {
		h.logger.Warn("validation failed: title is required")
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}

	if len(req.AuthorIDs) == 0 {
		h.logger.Warn("validation failed: author_ids is required")
		writeError(w, http.StatusBadRequest, "author_ids is required")
		return
	}

	h.logger.WithFields(logrus.Fields{
		"title":      req.Title,
		"author_ids": req.AuthorIDs,
	}).Debug("attempting to create book")

	book, err := h.store.CreateBook(r.Context(), req.Title, req.AuthorIDs)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrAuthorsMissing):
			h.logger.WithError(err).Warn("one or more authors do not exist")
			writeError(w, http.StatusBadRequest, "one or more authors do not exist")

		case errors.Is(err, store.ErrDuplicate):
			h.logger.WithError(err).Warn("book already exists")
			writeError(w, http.StatusConflict, "book already exists")

		default:
			h.logger.WithError(err).Error("failed to create book in store")
			writeError(w, http.StatusInternalServerError, "failed to create book")
		}
		return
	}

	h.logger.WithField("book_id", book.ID).Info("book successfully created")
	writeJSON(w, http.StatusCreated, book)
}
