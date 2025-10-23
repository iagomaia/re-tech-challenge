package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/iagomaia/re-tech-challenge/internal/infra/utils"
	"net/http"
	"os"
)

func GetDocsRoutes() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		utils.GetLogger().Info().Msg("Serving swagger.yaml")
		dir, _ := os.Getwd()
		if dir == "/" {
			dir = "/usr/local/bin"
		}
		filePath := dir + "/static/swagger.yaml"
		http.ServeFile(w, r, filePath)
	})
	return r
}
