package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/handler"
	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/store"
)

type Server struct {
	router *chi.Mux
}

func New(db *sql.DB) *Server {
	st := store.New(db)
	s := &Server{router: chi.NewRouter()}
	s.registerRoutes(handler.NewAuthors(st), handler.NewBooks(st))
	return s
}

func (s *Server) Listen() error {
	port := "3000"
	fmt.Printf("Server runs on port %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.router)
}

func (s *Server) Router() http.Handler {
	return s.router
}
