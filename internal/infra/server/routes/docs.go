package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"path"
	"strings"
)

func GetDocsRoutes() *chi.Mux {
	r := chi.NewRouter()
	dir, _ := os.Getwd()
	if dir == "/" {
		dir = "/usr/local/bin"
	}

	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		filePath := path.Join(dir, "static", "swagger.yaml")
		http.ServeFile(w, r, filePath)
	})

	swaggerUIDir := path.Join(dir, "static", "swagger-ui")
	fs := http.FileServer(http.Dir(swaggerUIDir))

	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, ".css") {
			w.Header().Set("Content-Type", "text/css")
		} else if strings.Contains(r.URL.Path, ".js") {
			w.Header().Set("Content-Type", "text/javascript")
		} else {
			w.Header().Set("Content-Type", "text/html")
		}

		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		http.StripPrefix(pathPrefix, fs).ServeHTTP(w, r)
	})
	return r
}
