package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootEndpoint(t *testing.T) {
	s := New(nil)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := "Hello World 123456"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rr.Body.String())
	}
}

func TestHealthEndpoint(t *testing.T) {
	s := New(nil)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := `{"status":"ok"}` + "\n"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rr.Body.String())
	}
}
