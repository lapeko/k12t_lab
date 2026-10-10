package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Api interface {
	Setup()
	Listen() error
}

type api struct {
	r *chi.Mux
}

func New() Api {
	return &api{
		r: chi.NewRouter(),
	}
}

func (a *api) Setup() {
	a.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World =)"))
	})
}

func (a *api) Listen() error {
	port := "3000"
	fmt.Printf("Server runs on port %s\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%s", port), a.r)
}
