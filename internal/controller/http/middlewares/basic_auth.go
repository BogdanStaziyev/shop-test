package middlewares

import (
	"net/http"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/responses"
)

// CheckAuth implements a simple middleware handler for adding basic http auth to a route.
func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			responses.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
			return
		} else if pass == "admin" && user == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		responses.ErrorResponse(w, http.StatusUnauthorized, "Unauthorized")
		return
	})
}
