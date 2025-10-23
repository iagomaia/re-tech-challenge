package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func GetHealthRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok"))
	})
	return r
}
