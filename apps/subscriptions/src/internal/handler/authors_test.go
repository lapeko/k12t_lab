package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthorsCreate_RequiresName(t *testing.T) {
	h := NewAuthors(nil)

	req := httptest.NewRequest(http.MethodPost, "/authors", strings.NewReader(`{"name":""}`))
	rr := httptest.NewRecorder()
	h.Create(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}
