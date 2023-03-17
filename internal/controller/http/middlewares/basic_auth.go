package middlewares

import (
	"net/http"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/responses"
)

// CheckAuth implements a simple middleware handler for adding basic http auth to a route.
func CheckAuth(password, user string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			us, pass, ok := r.BasicAuth()
			if !ok {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				responses.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
				return
			} else if pass == password && us == user {
				next.ServeHTTP(w, r)
				return
			}
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			responses.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		})
	}
}
