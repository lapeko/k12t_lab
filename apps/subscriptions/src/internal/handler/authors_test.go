package handler

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestAuthorsCreate_RequiresName(t *testing.T) {
	log := logrus.New()
	log.SetOutput(io.Discard)

	h := NewAuthors(nil, log)

	req := httptest.NewRequest(http.MethodPost, "/authors", strings.NewReader(`{"name":""}`))
	rr := httptest.NewRecorder()
	h.Create(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, rr.Code)
	}
}
