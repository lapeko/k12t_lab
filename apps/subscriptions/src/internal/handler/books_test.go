package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestBooksCreate_RequiresTitleAndAuthors(t *testing.T) {
	log := logrus.New()
	log.SetOutput(io.Discard)

	h := NewBooks(nil, log)

	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(`{"title":"","author_ids":[]}`))
	rr := httptest.NewRecorder()
	h.Create(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}
