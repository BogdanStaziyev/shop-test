package v1

import (
	"net/http"

	// External chi
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// Internal
	"github.com/BogdanStaziyev/shop-test/internal/controller/http/responses"
)

// Router initialize new CHI router
func Router(router *chi.Mux) http.Handler {
	router.Use(middleware.RedirectSlashes, middleware.Logger)

	router.Route("/api", func(apiRouter chi.Router) {
		// Health
		apiRouter.Route("/ping", func(healthRouter chi.Router) {
			healthRouter.Get("/", PingHandler())
			healthRouter.Handle("/*", NotFoundJSON())
		})
		apiRouter.Route("/v1", func(apiRouter chi.Router) {
			// Public routes
			apiRouter.Group(func(apiRouter chi.Router) {
				apiRouter.Handle("/*", NotFoundJSON())
			})
		})
	})
	return router
}

func NotFoundJSON() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.ErrorResponse(w, http.StatusNotFound, "Resource Not Found")
	}
}

func PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		responses.Response(w, http.StatusOK, "OK")
	}
}
