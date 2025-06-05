package routers

import (
	"cutURL/internal/handlers"
	"cutURL/internal/storage"
	"github.com/go-chi/chi/v5"
)

func RouterNew() *chi.Mux {
	r := chi.NewRouter()
	s := storage.NewStorage()

	r.Post("/", handlers.CreateUrlHandler(s))
	r.Get("/{id}", handlers.GetUrlHandler(s))

	return r
}
