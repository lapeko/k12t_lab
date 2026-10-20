package handler

import (
	"encoding/json"
	"net/http"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/model"
	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/store"
	"github.com/sirupsen/logrus"
)

type Authors struct {
	store  store.Author
	logger *logrus.Entry
}

func NewAuthors(store store.Author, logger *logrus.Logger) *Authors {
	return &Authors{
		store:  store,
		logger: logger.WithField("component", "authors_handler"),
	}
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
	var req model.CreateAuthorsRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	var names []string
	for _, name := range req.Names {
		if name == "" {
			h.logger.Errorf("create author error: property \"name\" shouldn't be empty\n")
		} else {
			names = append(names, name)
		}
	}

	authors, err := h.store.CreateAuthors(r.Context(), names)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "failed to create authors")
		return
	}

	writeJSON(w, http.StatusCreated, authors)
}
