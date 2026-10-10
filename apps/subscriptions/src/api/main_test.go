package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetup_RootEndpoint(t *testing.T) {
	a := New().(*api)
	a.Setup()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	a.r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, rr.Code)
	}

	expectedBody := "Hello World 1234"
	if rr.Body.String() != expectedBody {
		t.Errorf("expected body %q, got %q", expectedBody, rr.Body.String())
	}
}
