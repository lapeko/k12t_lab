package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/model"
	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/store"
)

type BookStore interface {
	ListBooks(ctx context.Context) ([]model.BookResponse, error)
	CreateBook(ctx context.Context, title string, authorIDs []int64) (model.Book, error)
}

type Books struct {
	store BookStore
}

func NewBooks(store BookStore) *Books {
	return &Books{store: store}
}

func (h *Books) List(w http.ResponseWriter, r *http.Request) {
	books, err := h.store.ListBooks(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list books")
		return
	}

	writeJSON(w, http.StatusOK, books)
}

func (h *Books) Create(w http.ResponseWriter, r *http.Request) {
	var req model.Book
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	if req.Title == "" {
		writeError(w, http.StatusBadRequest, "title is required")
		return
	}
	if len(req.AuthorIDs) == 0 {
		writeError(w, http.StatusBadRequest, "author_ids is required")
		return
	}

	book, err := h.store.CreateBook(r.Context(), req.Title, req.AuthorIDs)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrAuthorsMissing):
			writeError(w, http.StatusBadRequest, "one or more authors do not exist")
		case errors.Is(err, store.ErrDuplicate):
			writeError(w, http.StatusConflict, "book already exists")
		default:
			writeError(w, http.StatusInternalServerError, "failed to create book")
		}
		return
	}

	writeJSON(w, http.StatusCreated, book)
}
