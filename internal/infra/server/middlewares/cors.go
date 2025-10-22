package middlewares

import (
	"net/http"
	"os"

	"github.com/go-chi/cors"
	"github.com/iagomaia/re-tech-challenge/internal/infra/utils"
)

func GetCorsMiddleware() func(next http.Handler) http.Handler {
	svEnv := utils.GetStringValueOrDefault(os.Getenv("SV_ENV"), "dev")
	origins := []string{"*"}
	debug := svEnv == "local"
	if svEnv != "prod" {
		origins = []string{"*"}
	}
	return cors.Handler(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "X-Cid"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            debug,
		MaxAge:           300,
	})
}
