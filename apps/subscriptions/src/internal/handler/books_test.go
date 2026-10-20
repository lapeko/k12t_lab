package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBooksCreate_RequiresTitleAndAuthors(t *testing.T) {
	h := NewBooks(nil)

	req := httptest.NewRequest(http.MethodPost, "/books", strings.NewReader(`{"title":"","author_ids":[]}`))
	rr := httptest.NewRecorder()
	h.Create(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}
