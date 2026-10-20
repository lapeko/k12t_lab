package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/model"
)

type AuthorStore interface {
	ListAuthors(ctx context.Context) ([]model.Author, error)
	CreateAuthor(ctx context.Context, name string) (model.Author, error)
}

type Authors struct {
	store AuthorStore
}

func NewAuthors(store AuthorStore) *Authors {
	return &Authors{store: store}
}

func (h *Authors) List(w http.ResponseWriter, r *http.Request) {
	authors, err := h.store.ListAuthors(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to list authors")
		return
	}

	writeJSON(w, http.StatusOK, authors)
}

func (h *Authors) Create(w http.ResponseWriter, r *http.Request) {
	var req model.Author
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	if req.Name == "" {
		writeError(w, http.StatusBadRequest, "name is required")
		return
	}

	author, err := h.store.CreateAuthor(r.Context(), req.Name)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create author")
		return
	}

	writeJSON(w, http.StatusCreated, author)
}
