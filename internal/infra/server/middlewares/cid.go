package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/iagomaia/re-tech-challenge/internal/domain/models/utils"
)

func CidMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cid := uuid.New()
		ctx := context.WithValue(r.Context(), utils.CidContextKey, cid.String())
		w.Header().Add("X-Cid", cid.String())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
