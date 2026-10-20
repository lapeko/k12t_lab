package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/handler"
	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/logger"
	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/store"
	"github.com/lapeko/k12t_lab/apps/subscriptions/src/internal/utils"
)

type Server struct {
	router *chi.Mux
	logger *logrus.Logger
}

func New(db *sql.DB) *Server {
	var log *logrus.Logger
	if env, err := utils.ParseAppEnv(); err != nil {
		log.Fatalf("%e\n", err)
	} else {
		log = logger.New(env)
	}
	st := store.New(db, log)
	s := &Server{
		router: chi.NewRouter(),
		logger: log,
	}
	s.registerRoutes(
		handler.NewAuthors(st, log),
		handler.NewBooks(st, log),
	)
	return s
}

func (s *Server) Listen() error {
	port := "3000"
	s.logger.Info("Server runs on port %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), s.router)
}

func (s *Server) Router() http.Handler {
	return s.router
}
