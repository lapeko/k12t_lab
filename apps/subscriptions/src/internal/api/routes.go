package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/handler"
)

func (s *Server) registerRoutes(authors *handler.Authors, books *handler.Books) {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World 123456"))
	})

	s.router.Route("/authors", func(r chi.Router) {
		r.Get("/", authors.List)
		r.Post("/", authors.Create)
	})

	s.router.Route("/books", func(r chi.Router) {
		r.Get("/", books.List)
		r.Post("/", books.Create)
	})
}
